package dto

import (
	"github.com/horockey/poller_backend/internal/adapter/event"
	"github.com/horockey/poller_backend/internal/model"
)

type Msg struct {
	Poll   *Poll        `msgpack:"poll"`
	Action event.Action `msgpack:"action"`
}

func NewMsg(p *model.Poll, act event.Action) *Msg {
	return &Msg{
		Poll:   newPoll(p),
		Action: act,
	}
}
