package core

import (
	"github.com/TcMits/bookingdotcom/model"
)

const (
	PreMutateEventAction = "pre_mutate"
	PosMutateEventAction = "post_mutate"
	PostFindEventAction  = "post_find"
)

type DaoEvent struct {
	model  model.Model
	action string
}

func NewDaoEvent(model model.Model, action string) *DaoEvent {
	return &DaoEvent{
		model:  model,
		action: action,
	}
}

func (e *DaoEvent) Model() model.Model {
	return e.model
}

func (e *DaoEvent) Action() string {
	return e.action
}
