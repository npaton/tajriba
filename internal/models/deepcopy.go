// generated by deep-copy --pointer-receiver --type Group --type Link --type Scope --type Step --type Transition --type ParticipantChange --type StepChange --type AttributeChange --type ChangePayload -o deepcopy.go .; DO NOT EDIT.

package models

import (
	"time"
)

// DeepCopy generates a deep copy of *Group
func (o *Group) DeepCopy() *Group {
	var cp Group = *o
	if o.Links != nil {
		cp.Links = make([]*Link, len(o.Links))
		copy(cp.Links, o.Links)
		for i2 := range o.Links {
			if o.Links[i2] != nil {
				cp.Links[i2] = o.Links[i2].DeepCopy()
			}
		}
	}
	return &cp
}

// DeepCopy generates a deep copy of *Link
func (o *Link) DeepCopy() *Link {
	var cp Link = *o
	if o.Participant != nil {
		cp.Participant = o.Participant.DeepCopy()
	}
	return &cp
}

// DeepCopy generates a deep copy of *Scope
func (o *Scope) DeepCopy() *Scope {
	var cp Scope = *o
	if o.Name != nil {
		cp.Name = new(string)
		*cp.Name = *o.Name
	}
	if o.Kind != nil {
		cp.Kind = new(string)
		*cp.Kind = *o.Kind
	}
	if o.Attributes != nil {
		cp.Attributes = make([]*Attribute, len(o.Attributes))
		copy(cp.Attributes, o.Attributes)
		for i2 := range o.Attributes {
			if o.Attributes[i2] != nil {
				cp.Attributes[i2] = o.Attributes[i2].DeepCopy()
			}
		}
	}
	if o.AttributesMap != nil {
		cp.AttributesMap = make(map[string]*Attribute, len(o.AttributesMap))
		for k2, v2 := range o.AttributesMap {
			var cp_AttributesMap_v2 *Attribute
			if v2 != nil {
				cp_AttributesMap_v2 = v2.DeepCopy()
			}
			cp.AttributesMap[k2] = cp_AttributesMap_v2
		}
	}
	if o.Links != nil {
		cp.Links = make([]*Link, len(o.Links))
		copy(cp.Links, o.Links)
		for i2 := range o.Links {
			if o.Links[i2] != nil {
				cp.Links[i2] = o.Links[i2].DeepCopy()
			}
		}
	}
	return &cp
}

// DeepCopy generates a deep copy of *Step
func (o *Step) DeepCopy() *Step {
	var cp Step = *o
	if o.StartedAt != nil {
		cp.StartedAt = new(time.Time)
		*cp.StartedAt = *o.StartedAt
	}
	if o.EndedAt != nil {
		cp.EndedAt = new(time.Time)
		*cp.EndedAt = *o.EndedAt
	}
	if o.Transitions != nil {
		cp.Transitions = make([]*Transition, len(o.Transitions))
		copy(cp.Transitions, o.Transitions)
		for i2 := range o.Transitions {
			if o.Transitions[i2] != nil {
				cp.Transitions[i2] = o.Transitions[i2].DeepCopy()
			}
		}
	}
	if o.Links != nil {
		cp.Links = make([]*Link, len(o.Links))
		copy(cp.Links, o.Links)
		for i2 := range o.Links {
			if o.Links[i2] != nil {
				cp.Links[i2] = o.Links[i2].DeepCopy()
			}
		}
	}
	return &cp
}

// DeepCopy generates a deep copy of *Transition
func (o *Transition) DeepCopy() *Transition {
	var cp Transition = *o
	return &cp
}

// DeepCopy generates a deep copy of *ParticipantChange
func (o *ParticipantChange) DeepCopy() *ParticipantChange {
	var cp ParticipantChange = *o
	return &cp
}

// DeepCopy generates a deep copy of *StepChange
func (o *StepChange) DeepCopy() *StepChange {
	var cp StepChange = *o
	if o.Since != nil {
		cp.Since = new(time.Time)
		*cp.Since = *o.Since
	}
	if o.Remaining != nil {
		cp.Remaining = new(int)
		*cp.Remaining = *o.Remaining
	}
	if o.Ellapsed != nil {
		cp.Ellapsed = new(int)
		*cp.Ellapsed = *o.Ellapsed
	}
	return &cp
}

// DeepCopy generates a deep copy of *AttributeChange
func (o *AttributeChange) DeepCopy() *AttributeChange {
	var cp AttributeChange = *o
	if o.Index != nil {
		cp.Index = new(int)
		*cp.Index = *o.Index
	}
	if o.Val != nil {
		cp.Val = new(string)
		*cp.Val = *o.Val
	}
	return &cp
}

// DeepCopy generates a deep copy of *ChangePayload
func (o *ChangePayload) DeepCopy() *ChangePayload {
	var cp ChangePayload = *o
	return &cp
}
