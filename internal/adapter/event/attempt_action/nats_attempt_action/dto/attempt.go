package dto

import (
	"time"

	"github.com/horockey/poller_backend/internal/model"
)

type Attempt struct {
	ID          string           `msgpack:"id"`
	TimeCreated time.Time        `msgpack:"time_created"`
	PollID      string           `msgpack:"poll_id"`
	Answers     []*AttemptAnswer `msgpack:"answers"`
}

func newAtempt(a *model.Attempt) *Attempt {
	answers := make([]*AttemptAnswer, 0, len(a.Answers))
	for _, ans := range a.Answers {
		answers = append(answers, newAttemptAnswer(ans))
	}

	return &Attempt{
		ID:          a.ID,
		TimeCreated: a.TimeCreated,
		PollID:      a.PollID,
		Answers:     answers,
	}
}
