package dto

import (
	"github.com/horockey/poller_backend/internal/adapter/event"
	"github.com/horockey/poller_backend/internal/model"
)

type Msg struct {
	Attempt *Attempt     `msgpack:"attempt"`
	Action  event.Action `msgpack:"action"`
}

func NewMsg(a *model.Attempt, act event.Action) *Msg {
	return &Msg{
		Attempt: newAtempt(a),
		Action:  act,
	}
}
