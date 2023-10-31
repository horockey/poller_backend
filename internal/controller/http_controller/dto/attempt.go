package dto

import (
	"time"

	"github.com/horockey/poller_backend/internal/model"
)

type Attempt struct {
	ID      string           `json:"id"`
	Ts      string           `json:"ts"`
	PollID  string           `json:"poll_id"`
	Answers []*AttemptAnswer `json:"answers"`
}

type AttemptAnswer struct {
	Question *Question `json:"question"`
	Answers  []*Answer `json:"answers"`
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
		Ts:      a.Ts.Format(time.RFC3339),
		PollID:  a.PollID,
		Answers: answers,
	}
}

func newAttemptAnswer(a *model.AttemptAnswer) *AttemptAnswer {
	answers := make([]*Answer, len(a.Answers))
	for _, ans := range a.Answers {
		answers = append(answers, NewAnswer(ans))
	}

	return &AttemptAnswer{
		Question: NewQuestion(a.Question),
		Answers:  answers,
	}
}
