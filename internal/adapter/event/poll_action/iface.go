package poll_action

import (
	"github.com/horockey/poller_backend/internal/adapter/event"
	"github.com/horockey/poller_backend/internal/model"
)

type Event interface {
	Send(p *model.Poll, act event.Action) error
}
