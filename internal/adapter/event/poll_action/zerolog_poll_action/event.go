package zerolog_poll_action

import (
	"encoding/json"
	"fmt"

	"github.com/horockey/poller_backend/internal/adapter/event"
	"github.com/horockey/poller_backend/internal/adapter/event/poll_action"
	"github.com/horockey/poller_backend/internal/model"
	"github.com/rs/zerolog"
)

var _ poll_action.Event = &zerologPollAction{}

type zerologPollAction struct {
	logger zerolog.Logger
}

func New(logger zerolog.Logger) *zerologPollAction {
	return &zerologPollAction{
		logger: logger,
	}
}

func (ev *zerologPollAction) Send(a *model.Poll, act event.Action) error {
	data, err := json.Marshal(a)
	if err != nil {
		return fmt.Errorf("marshalling attempt: %w", err)
	}

	ev.logger.Info().
		Str("action", act.String()).
		RawJSON("attempt", data).
		Send()

	return nil
}
