package zerolog_attempt_action

import (
	"encoding/json"
	"fmt"

	"github.com/horockey/poller_backend/internal/adapter/event"
	"github.com/horockey/poller_backend/internal/adapter/event/attempt_action"
	"github.com/horockey/poller_backend/internal/model"
	"github.com/rs/zerolog"
)

var _ attempt_action.Event = &zerologAttemptAction{}

type zerologAttemptAction struct {
	logger zerolog.Logger
}

func New(logger zerolog.Logger) *zerologAttemptAction {
	return &zerologAttemptAction{
		logger: logger,
	}
}

func (ev *zerologAttemptAction) Send(a *model.Attempt, act event.Action) error {
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
