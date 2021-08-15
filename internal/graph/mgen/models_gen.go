// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package mgen

import (
	"fmt"
	"io"
	"strconv"

	"github.com/empiricaly/tajriba/internal/models"
)

type Change interface {
	IsChange()
}

// AddGroupInput creates a new Group.
type AddGroupInput struct {
	// participantIDs are the IDs of the Participants to link with the Group.
	ParticipantIDs []string `json:"participantIDs"`
}

// AddGroupPayload is the return payload for the addGroup mutation.
type AddGroupPayload struct {
	// group that the participant is added to.
	Group *models.Group `json:"group"`
}

// AddParticipantInput finds or creates a Participant by Unique Identifier.
type AddParticipantInput struct {
	// identifier is the unique identifier for the Pariticipant. The identifier
	// is how a participant "logs into" the system.
	Identifier string `json:"identifier"`
}

// AddParticipantPayload is the return payload for the addParticipant mutation.
type AddParticipantPayload struct {
	// participant is the created Participants.
	Participant *models.Participant `json:"participant"`
	// sessionToken is the session token to be used for authenticated requets.
	SessionToken string `json:"sessionToken"`
}

// AddScopeInput creates a new Scope.
type AddScopeInput struct {
	// name is the *unique* name of the Scope. If a scope with the same name already
	// exists, it will return an "already exists" error.
	Name string `json:"name"`
	// attributes to be attached to the Scope at creation.
	Attributes []*SetAttributeInput `json:"attributes"`
}

// AddScopePayload is the return payload for the addScope mutation.
type AddScopePayload struct {
	// scope that the participant is added to.
	Scope *models.Scope `json:"scope"`
}

// AddStepInput creates a new Step.
type AddStepInput struct {
	// duration is the duration in seconds of the Step should last before ending.
	Duration int `json:"duration"`
}

// AddStepPayload is the return payload for the addStep mutation.
type AddStepPayload struct {
	// step that the participant is added to.
	Step *models.Step `json:"step"`
}

type AttributeConnection struct {
	TotalCount int              `json:"totalCount"`
	PageInfo   *PageInfo        `json:"pageInfo"`
	Edges      []*AttributeEdge `json:"edges"`
}

type AttributeEdge struct {
	Node   *models.Attribute `json:"node"`
	Cursor string            `json:"cursor"`
}

type ChangePayload struct {
	// change is the Change.
	Change Change `json:"change"`
	// removed indicates whether the record was removed from scope.
	Removed bool `json:"removed"`
	// done indicates that the state has finished synchorizing.
	Done bool `json:"done"`
}

type GroupConnection struct {
	TotalCount int          `json:"totalCount"`
	PageInfo   *PageInfo    `json:"pageInfo"`
	Edges      []*GroupEdge `json:"edges"`
}

type GroupEdge struct {
	Node   *models.Group `json:"node"`
	Cursor string        `json:"cursor"`
}

type LinkConnection struct {
	TotalCount int         `json:"totalCount"`
	PageInfo   *PageInfo   `json:"pageInfo"`
	Edges      []*LinkEdge `json:"edges"`
}

type LinkEdge struct {
	Node   *models.Link `json:"node"`
	Cursor string       `json:"cursor"`
}

// LinkInput links or unlinks Participants with a Node.
type LinkInput struct {
	// nodeIDs are the IDs of the Nodes that the Participants should be added to.
	NodeIDs []string `json:"nodeIDs"`
	// participantIDs are the IDs of the Participants that should be added to the
	// Nodes.
	ParticipantIDs []string `json:"participantIDs"`
	// link indicates whether the Participant was linked or unlinked with the Node.
	// WARNING: UNLINKING NOT CURRENTLY SUPPORTED, link must be true.
	Link bool `json:"link"`
}

// LinkPayload is the return payload for the assignParticipants
// mutation.
type LinkPayload struct {
	// nodes the participants are added to.
	Nodes []models.Node `json:"nodes"`
	// participants that are added to the Node.
	Participants []*models.Participant `json:"participants"`
}

// LoginInput is the input for login()
type LoginInput struct {
	// username is the user identifier string.
	Username string `json:"username"`
	// password of the user.
	Password string `json:"password"`
}

// LoginPayload is returned by login()
type LoginPayload struct {
	// user is the logged in User.
	User *models.User `json:"user"`
	// sessionToken is the session token to be used for authenticated requets.
	SessionToken string `json:"sessionToken"`
}

// OnAnyEventInput is the input for the onAnyEvent subscription.
type OnAnyEventInput struct {
	// nodeID is an optional node ID of the node to listen to. If nodeID is
	// specified, nodeType must also be given.
	NodeID *string `json:"nodeID"`
}

// OnEventInput is the input for the onEvent subscription.
type OnEventInput struct {
	// eventsTypes speficies which events to listen to.
	EventTypes []EventType `json:"eventTypes"`
	// nodeID is an optional node ID of the node to listen to. If nodeID is
	// specified, nodeType must also be given.
	NodeID *string `json:"nodeID"`
}

// OnEventPayload is the payload for the on[Any]Event subscriptions.
type OnEventPayload struct {
	// eventID is a unique identifier for the current event. Each OnEventPayload for
	// a single client will have a different eventID. eventID will be the same for
	// different clients on the same event.
	EventID string `json:"eventID"`
	// eventType is the type of the current event.
	EventType EventType `json:"eventType"`
	// node is the node that triggered the event
	Node models.Node `json:"node"`
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

type ParticipantConnection struct {
	TotalCount int                `json:"totalCount"`
	PageInfo   *PageInfo          `json:"pageInfo"`
	Edges      []*ParticipantEdge `json:"edges"`
}

type ParticipantEdge struct {
	Node   *models.Participant `json:"node"`
	Cursor string              `json:"cursor"`
}

// RegisterServiceInput is the input for registerService()
type RegisterServiceInput struct {
	// name is the name of the service to register.
	Name string `json:"name"`
	// token is the Service Registration token.
	Token string `json:"token"`
}

// RegisterServicePayload is returned by registerService()
type RegisterServicePayload struct {
	// service is newly created Service.
	Service *models.Service `json:"service"`
	// sessionToken is the session token to be used for authenticated requets.
	SessionToken string `json:"sessionToken"`
}

type ScopeConnection struct {
	TotalCount int          `json:"totalCount"`
	PageInfo   *PageInfo    `json:"pageInfo"`
	Edges      []*ScopeEdge `json:"edges"`
}

type ScopeEdge struct {
	Node   *models.Scope `json:"node"`
	Cursor string        `json:"cursor"`
}

// ScopedAttributesPayload is the return payload for the addScope mutation.
type ScopedAttributesPayload struct {
	// scope that the participant is added to. Attribute may be null only if the
	// subscription did not match any Scopes and done must be published.
	Attribute *models.Attribute `json:"attribute"`
	// done indicates that the state has finished synchorizing.
	Done bool `json:"done"`
	// isNew returns true if the Attribute for key and nodeID was just created.
	IsNew bool `json:"isNew"`
}

// SetAttributeInput sets an Attribute on a Node.
type SetAttributeInput struct {
	// key identifies the unique key of the Attribute.
	Key string `json:"key"`
	// val is the value of the Attribute. It can be any JSON encodable value. If
	// value is not defined, value is assumed to be `null`.
	Val *string `json:"val"`
	// index of value if Attribute is a vector. An Attribute cannot mutate between
	// vector and non-vector formats.
	Index *int `json:"index"`
	// vector indicates the Attribute is a vector.
	Vector *bool `json:"vector"`
	// append allows appending to a vector without specifying the index.
	Append *bool `json:"append"`
	// private indicates whether the Attribute shouldn't be visible to Participants
	// in the scope.
	// private must be set on the Attribute at creation.
	// Defaults to false and does need to be sent on subsequent updates.
	Private *bool `json:"private"`
	// protected indicates the Attribute cannot be modified by other Participants. A
	// Participant can only set protected Records on their Participant record.
	// Users and Services can update protected Attributes.
	// protected must be set on the Attribute at creation.
	// Defaults to false and does need to be sent on subsequent updates.
	Protected *bool `json:"protected"`
	// immutable indicates the Attribute can never be changed by any Actor.
	// immutable must be set on the Attribute at creation.
	// Defaults to false and does need to be sent on subsequent updates.
	Immutable *bool `json:"immutable"`
	// ID of object on which to update the value. NodeID is required if attribute is
	// not created with addScope().
	NodeID *string `json:"nodeID"`
}

// SetAttributePayload is the return payload for the setAttribute mutation.
type SetAttributePayload struct {
	// attribute is the Attribute updated.
	Attribute *models.Attribute `json:"attribute"`
}

type StepConnection struct {
	TotalCount int         `json:"totalCount"`
	PageInfo   *PageInfo   `json:"pageInfo"`
	Edges      []*StepEdge `json:"edges"`
}

type StepEdge struct {
	Node   *models.Step `json:"node"`
	Cursor string       `json:"cursor"`
}

type StepOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *StepOrderField `json:"field"`
}

type TransitionConnection struct {
	TotalCount int               `json:"totalCount"`
	PageInfo   *PageInfo         `json:"pageInfo"`
	Edges      []*TransitionEdge `json:"edges"`
}

type TransitionEdge struct {
	Node   *models.Transition `json:"node"`
	Cursor string             `json:"cursor"`
}

// TransitionInput transitions a Node.
type TransitionInput struct {
	// nodeID is the ID of the Node to transition.
	NodeID string `json:"nodeID"`
	// from is the current State of the Node. To avoid concurrency or repeat errors,
	// from is required, and the transition will not happen if the from State does
	// not correspond to the Node's current State.
	From models.State `json:"from"`
	// to is the desired State of the Step.
	To models.State `json:"to"`
	// cause is an optional open string explaining the reason for the transition.
	Cause *string `json:"cause"`
}

// TransitionPayload is the return payload for the transition() mutation.
type TransitionPayload struct {
	// transition is the created Transition.
	Transition *models.Transition `json:"transition"`
}

// EventType holds types of event that can trigger hooks.
type EventType string

const (
	EventTypeStepAdd               EventType = "STEP_ADD"
	EventTypeScopeAdd              EventType = "SCOPE_ADD"
	EventTypeGroupAdd              EventType = "GROUP_ADD"
	EventTypeTransitionAdd         EventType = "TRANSITION_ADD"
	EventTypeParticipantAdd        EventType = "PARTICIPANT_ADD"
	EventTypeParticipantConnect    EventType = "PARTICIPANT_CONNECT"
	EventTypeParticipantDisconnect EventType = "PARTICIPANT_DISCONNECT"
	EventTypeLinkAdd               EventType = "LINK_ADD"
	EventTypeAttributeUpdate       EventType = "ATTRIBUTE_UPDATE"
)

var AllEventType = []EventType{
	EventTypeStepAdd,
	EventTypeScopeAdd,
	EventTypeGroupAdd,
	EventTypeTransitionAdd,
	EventTypeParticipantAdd,
	EventTypeParticipantConnect,
	EventTypeParticipantDisconnect,
	EventTypeLinkAdd,
	EventTypeAttributeUpdate,
}

func (e EventType) IsValid() bool {
	switch e {
	case EventTypeStepAdd, EventTypeScopeAdd, EventTypeGroupAdd, EventTypeTransitionAdd, EventTypeParticipantAdd, EventTypeParticipantConnect, EventTypeParticipantDisconnect, EventTypeLinkAdd, EventTypeAttributeUpdate:
		return true
	}
	return false
}

func (e EventType) String() string {
	return string(e)
}

func (e *EventType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventType", str)
	}
	return nil
}

func (e EventType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderDirection string

const (
	OrderDirectionAsc  OrderDirection = "ASC"
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Role string

const (
	// ADMIN is priviledged access for Users and Services.
	RoleAdmin Role = "ADMIN"
	// PARTICIPANT is access tailored for Participants' needs.
	RoleParticipant Role = "PARTICIPANT"
)

var AllRole = []Role{
	RoleAdmin,
	RoleParticipant,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleParticipant:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StepOrderField string

const (
	StepOrderFieldCreatedAt StepOrderField = "CREATED_AT"
	StepOrderFieldStartedAt StepOrderField = "STARTED_AT"
	StepOrderFieldDuration  StepOrderField = "DURATION"
)

var AllStepOrderField = []StepOrderField{
	StepOrderFieldCreatedAt,
	StepOrderFieldStartedAt,
	StepOrderFieldDuration,
}

func (e StepOrderField) IsValid() bool {
	switch e {
	case StepOrderFieldCreatedAt, StepOrderFieldStartedAt, StepOrderFieldDuration:
		return true
	}
	return false
}

func (e StepOrderField) String() string {
	return string(e)
}

func (e *StepOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StepOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StepOrderField", str)
	}
	return nil
}

func (e StepOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
