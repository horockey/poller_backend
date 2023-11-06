package nats_poll_action

import (
	"fmt"

	"github.com/horockey/poller_backend/internal/adapter/event"
	"github.com/horockey/poller_backend/internal/adapter/event/poll_action"
	"github.com/horockey/poller_backend/internal/adapter/event/poll_action/nats_poll_action/dto"
	"github.com/horockey/poller_backend/internal/model"
	"github.com/nats-io/nats.go"
	"github.com/vmihailenco/msgpack"
)

var _ poll_action.Event = &natsPollAction{}

type natsPollAction struct {
	conn  *nats.Conn
	topic string
}

func New(conn *nats.Conn, topic string) *natsPollAction {
	return &natsPollAction{
		conn:  conn,
		topic: topic,
	}
}

func (ev *natsPollAction) Send(p *model.Poll, act event.Action) error {
	data, err := msgpack.Marshal(dto.NewMsg(p, act))
	if err != nil {
		return fmt.Errorf("marshaling msg: %w", err)
	}

	if err := ev.conn.Publish(ev.topic, data); err != nil {
		return fmt.Errorf("publishing message: %w", err)
	}

	return nil
}
