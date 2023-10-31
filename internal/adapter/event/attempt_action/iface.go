package attempt_action

import (
	"github.com/horockey/poller_backend/internal/adapter/event"
	"github.com/horockey/poller_backend/internal/model"
)

type Event interface {
	Send(a *model.Attempt, act event.Action) error
}
