package dto

import (
	"fmt"
	"time"

	"github.com/horockey/poller_backend/internal/model"
)

type Attempt struct {
	ID      string           `json:"id"`
	Ts      string           `json:"ts"`
	PollID  string           `json:"poll_id"`
	Answers []*AttemptAnswer `json:"answers"`
}

func NewAttempts(ats []*model.Attempt) []*Attempt {
	res := make([]*Attempt, len(ats))

	for _, at := range ats {
		res = append(res, NewAttempt(at))
	}

	return res
}

func NewAttempt(a *model.Attempt) *Attempt {
	answers := make([]*AttemptAnswer, len(a.Answers))
	for _, aa := range a.Answers {
		answers = append(answers, newAttemptAnswer(aa))
	}

	return &Attempt{
		ID:      a.ID,
		Ts:      a.TimeCreated.Format(time.RFC3339),
		PollID:  a.PollID,
		Answers: answers,
	}
}

func (a *Attempt) ToModel() (*model.Attempt, error) {
	ts, err := time.Parse(time.RFC3339, a.Ts)
	if err != nil {
		return nil, fmt.Errorf("parsing ts: %w", err)
	}

	answers := make([]*model.AttemptAnswer, len(a.Answers))
	for _, ans := range a.Answers {
		answers = append(answers, ans.ToModel())
	}

	return &model.Attempt{
		ID:          a.ID,
		TimeCreated: ts,
		PollID:      a.PollID,
		Answers:     answers,
	}, nil
}
