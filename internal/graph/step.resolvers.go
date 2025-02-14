package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/empiricaly/tajriba/internal/graph/mgen"
	"github.com/empiricaly/tajriba/internal/models"
	"github.com/empiricaly/tajriba/internal/runtime"
	errs "github.com/pkg/errors"
)

// AddSteps is the resolver for the addSteps field.
func (r *mutationResolver) AddSteps(ctx context.Context, input []*mgen.AddStepInput) ([]*mgen.AddStepPayload, error) {
	rt := runtime.ForContext(ctx)

	p := make([]*mgen.AddStepPayload, len(input))

	for i, sinput := range input {
		s, err := rt.AddStep(ctx, sinput.Duration)
		if err != nil {
			return nil, errs.Wrap(err, "add step")
		}

		p[i] = &mgen.AddStepPayload{Step: s}
	}

	return p, nil
}

// Steps is the resolver for the steps field.
func (r *queryResolver) Steps(ctx context.Context, after *string, first *int, before *string, last *int) (*mgen.StepConnection, error) {
	rt := runtime.ForContext(ctx)

	scopes, total, hasNext, hasPrev, err := rt.Steps(ctx, after, first, before, last)
	if err != nil {
		return nil, errs.Wrap(err, "get scopes")
	}

	var start, end string

	l := len(scopes)
	edges := make([]*mgen.StepEdge, l)

	for i, scope := range scopes {
		edges[i] = &mgen.StepEdge{
			Node:   scope,
			Cursor: scope.ID,
		}

		if i == 0 {
			start = scope.ID
		}

		if i == l-1 {
			end = scope.ID
		}
	}

	return &mgen.StepConnection{
		TotalCount: total,
		PageInfo: &mgen.PageInfo{
			HasNextPage:     hasNext,
			HasPreviousPage: hasPrev,
			StartCursor:     &start,
			EndCursor:       &end,
		},
		Edges: edges,
	}, nil
}

// Transitions is the resolver for the transitions field.
func (r *stepResolver) Transitions(ctx context.Context, obj *models.Step, after *string, first *int, before *string, last *int) (*mgen.TransitionConnection, error) {
	rt := runtime.ForContext(ctx)

	transitions, total, hasNext, hasPrev, err := rt.StepTransitions(ctx, obj, after, first, before, last)
	if err != nil {
		return nil, errs.Wrap(err, "get step participations")
	}

	var start, end string

	l := len(transitions)
	edges := make([]*mgen.TransitionEdge, l)

	for i, transition := range transitions {
		edges[i] = &mgen.TransitionEdge{
			Node:   transition,
			Cursor: transition.ID,
		}

		if i == 0 {
			start = transition.ID
		}

		if i == l-1 {
			end = transition.ID
		}
	}

	return &mgen.TransitionConnection{
		TotalCount: total,
		PageInfo: &mgen.PageInfo{
			HasNextPage:     hasNext,
			HasPreviousPage: hasPrev,
			StartCursor:     &start,
			EndCursor:       &end,
		},
		Edges: edges,
	}, nil
}

// Links is the resolver for the links field.
func (r *stepResolver) Links(ctx context.Context, obj *models.Step, after *string, first *int, before *string, last *int) (*mgen.LinkConnection, error) {
	rt := runtime.ForContext(ctx)

	links, total, hasNext, hasPrev, err := rt.StepLinks(ctx, obj, after, first, before, last)
	if err != nil {
		return nil, errs.Wrap(err, "get step links")
	}

	var start, end string

	l := len(links)
	edges := make([]*mgen.LinkEdge, l)

	for i, link := range links {
		edges[i] = &mgen.LinkEdge{
			Node:   link,
			Cursor: link.ID,
		}

		if i == 0 {
			start = link.ID
		}

		if i == l-1 {
			end = link.ID
		}
	}

	return &mgen.LinkConnection{
		TotalCount: total,
		PageInfo: &mgen.PageInfo{
			HasNextPage:     hasNext,
			HasPreviousPage: hasPrev,
			StartCursor:     &start,
			EndCursor:       &end,
		},
		Edges: edges,
	}, nil
}

// Step returns StepResolver implementation.
func (r *Resolver) Step() StepResolver { return &stepResolver{r} }

type stepResolver struct{ *Resolver }
