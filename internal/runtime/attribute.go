package runtime

import (
	"context"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/empiricaly/tajriba/internal/auth/actor"
	"github.com/empiricaly/tajriba/internal/graph/mgen"
	"github.com/empiricaly/tajriba/internal/models"
	"github.com/empiricaly/tajriba/internal/server/metadata"
	"github.com/empiricaly/tajriba/internal/store"
	"github.com/empiricaly/tajriba/internal/utils/ids"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (r *Runtime) SetAttributes(
	ctx context.Context,
	inputs []*mgen.SetAttributeInput,
) (attrs []*models.Attribute, err error) {
	if len(inputs) == 0 {
		return nil, ErrEmptyInput
	}

	r.Lock()
	defer r.Unlock()

	attrs, err = r.prepAttributes(ctx, inputs)
	if err != nil {
		return nil, errors.Wrap(err, "check attributes")
	}

	attrs, err = r.setAttributes(ctx, attrs)
	if err != nil {
		return nil, errors.Wrap(err, "set attributes")
	}

	return attrs, nil
}

func (r *Runtime) prepAttributes(
	ctx context.Context,
	inputs []*mgen.SetAttributeInput,
) ([]*models.Attribute, error) {
	actr := actor.ForContext(ctx)
	if actr == nil {
		md := metadata.RequestForContext(ctx)
		spew.Dump(md)
		return nil, ErrNotAuthorized
	}

	now := time.Now()
	actorID := actr.GetID()
	attrs := make([]*models.Attribute, 0, len(inputs))

	for _, input := range inputs {
		if input.NodeID == nil {
			return nil, errors.New("nodeID is required")
		}

		scope, ok := r.scopesMap[*input.NodeID]
		if !ok {
			return nil, errors.New("scope not found")
		}

		if scope.Name != nil && *scope.Name == "global" {
			if _, ok := actr.(*models.Participant); ok {
				return nil, ErrNotAuthorized
			}
		}

		index := input.Index
		var vector, private, protected, immutable bool
		if input.Append != nil && *input.Append {
			vector = true

			var maxIndex int
			var found bool
			for _, attr := range scope.Attributes {
				if attr.Key == input.Key {
					if !attr.Vector {
						return nil, errors.Errorf("attribute for '%s' is not a vector", input.Key)
					}

					if attr.Index == nil {
						return nil, errors.Errorf("attribute for '%s' has no index", input.Key)
					}

					found = true
					if *attr.Index > maxIndex {
						maxIndex = *attr.Index
					}
				}
			}

			if !found {
				index = &maxIndex
			} else {
				maxIndex++
				index = &maxIndex
			}
		} else if input.Index != nil {
			vector = true
		}

		if input.Private != nil && *input.Private {
			private = true
		}

		if input.Protected != nil && *input.Protected {
			protected = true
		}

		if input.Immutable != nil && *input.Immutable {
			immutable = true
		}

		a := &models.Attribute{
			ID:          ids.ID(ctx),
			CreatedAt:   now,
			CreatedBy:   actr,
			CreatedByID: actorID,
			Key:         input.Key,
			Val:         input.Val,
			Index:       index,
			Vector:      vector,
			NodeID:      scope.ID,
			Node:        scope,
			Version:     1,
			Private:     private,
			Protected:   protected,
			Immutable:   immutable,
		}

		last := scope.AttributesMap[a.LookupKey()]

		if last != nil {
			if last.Immutable {
				return nil, ErrImmutable
			}

			if (a.Vector && !last.Vector) || !a.Vector && last.Vector {
				return nil, errors.New("cannot mutate vector")
			}

			if last.Private {
				a.Private = true
			}

			if last.Protected {
				a.Protected = true
			}

			// Check protected
			if a.Protected {
				if _, ok := actr.(*models.Participant); ok {
					if last.CreatedByID != a.CreatedByID {
						return nil, ErrNotAuthorized
					}
				}
			}

			a.Version = last.Version + 1
			a.Versions = append(last.Versions, last)
		}

		attrs = append(attrs, a)
	}

	return attrs, nil
}

func (r *Runtime) setAttributes(
	ctx context.Context,
	attrs []*models.Attribute,
) (attributes []*models.Attribute, err error) {
	conn := store.ForContext(ctx)
	for _, attr := range attrs {
		err = conn.Save(attr)
		if err != nil {
			log.Ctx(r.ctx).Error().Err(err).Msg("runtime: failed to save attribute")

			continue
		}

		scope, _ := attr.Node.(*models.Scope)

		// Remove old version of attribute
		if len(attr.Versions) > 0 {
			var n, check int

			for _, a := range scope.Attributes {
				if a.LookupKey() != attr.LookupKey() {
					scope.Attributes[n] = a
					n++

					continue
				}

				check++

				if check > 1 {
					log.Ctx(r.ctx).Warn().
						Interface("attr1", a).
						Interface("attr2", attr).
						Msg("double attribute!")
				}
			}

			// Add new version to end of list
			scope.Attributes[len(scope.Attributes)-1] = attr
		} else {
			// Append new attribute to end of list
			scope.Attributes = append(scope.Attributes, attr)
		}

		scope.AttributesMap[attr.LookupKey()] = attr

		// Add attribute to global vars
		r.attributes = append(r.attributes, attr)
		r.attributesMap[attr.ID] = attr
		r.values[attr.ID] = attr

		r.propagateHook(ctx, mgen.EventTypeAttributeUpdate, attr.ID, attr)
	}

	if err := r.pushAttributes(ctx, attrs); err != nil {
		log.Ctx(r.ctx).Error().Err(err).Msg("runtime: failed to push new attributes to participants")
	}

	return attrs, nil
}

func (r *Runtime) AttributeVersions(
	ctx context.Context,
	attrID string,
	after *string,
	first *int,
	before *string,
	last *int,
) (
	attrs []*models.Attribute,
	total int,
	hasNext,
	hasPrev bool,
	err error,
) {
	r.RLock()
	defer r.RUnlock()

	attr, ok := r.attributesMap[attrID]
	if !ok {
		return nil, 0, false, false, ErrNotFound
	}

	items := make([]models.Cursorer, len(attr.Versions))
	for i := range attr.Versions {
		items[i] = attr.Versions[i]
	}

	items, total, hasNext, hasPrev, err = paginate(items, after, first, before, last)

	versions := make([]*models.Attribute, len(items))
	for i := range items {
		versions[i], _ = items[i].(*models.Attribute)
	}

	return versions, total, hasNext, hasPrev, err
}

func (r *Runtime) IsCurrent(
	ctx context.Context,
	attr *models.Attribute,
) (
	isCurrent bool,
	err error,
) {
	r.RLock()
	defer r.RUnlock()

	scope, ok := r.scopesMap[attr.NodeID]
	if !ok {
		return false, ErrNotFound
	}

	last, ok := scope.AttributesMap[attr.LookupKey()]
	if !ok {
		return false, ErrNotFound
	}

	return last.ID == attr.ID, nil
}

func (r *Runtime) pushAttributes(ctx context.Context, attrs []*models.Attribute) error {
	if err := r.pushAttributesForChanges(ctx, attrs); err != nil {
		return errors.Wrap(err, "push attributes for changes")
	}

	if err := r.pushAttributesForScopedAttributes(ctx, attrs); err != nil {
		return errors.Wrap(err, "push attributes for scoped attributes")
	}

	return nil
}
