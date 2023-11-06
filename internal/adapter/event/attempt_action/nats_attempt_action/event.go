package nats_attempt_action

import (
	"fmt"

	"github.com/horockey/poller_backend/internal/adapter/event"
	"github.com/horockey/poller_backend/internal/adapter/event/attempt_action"
	"github.com/horockey/poller_backend/internal/adapter/event/attempt_action/nats_attempt_action/dto"
	"github.com/horockey/poller_backend/internal/model"
	"github.com/nats-io/nats.go"
	"github.com/vmihailenco/msgpack"
)

var _ attempt_action.Event = &natsAttemptAction{}

type natsAttemptAction struct {
	conn  *nats.Conn
	topic string
}

func New(conn *nats.Conn, topic string) *natsAttemptAction {
	return &natsAttemptAction{
		conn:  conn,
		topic: topic,
	}
}

func (ev *natsAttemptAction) Send(a *model.Attempt, act event.Action) error {
	data, err := msgpack.Marshal(dto.NewMsg(a, act))
	if err != nil {
		return fmt.Errorf("marshaling msg: %w", err)
	}

	if err := ev.conn.Publish(ev.topic, data); err != nil {
		return fmt.Errorf("publishing message: %w", err)
	}

	return nil
}
