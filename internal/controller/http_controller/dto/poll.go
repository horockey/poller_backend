package dto

import (
	"fmt"
	"time"

	"github.com/horockey/poller_backend/internal/model"
)

type Poll struct {
	ID          string      `json:"id"`
	TimeCreated string      `json:"time_created"`
	Title       string      `json:"title"`
	Questions   []*Question `json:"questions"`
}

func NewPoll(p *model.Poll) *Poll {
	questions := make([]*Question, len(p.Questions))
	for _, q := range p.Questions {
		questions = append(questions, NewQuestion(q))
	}

	return &Poll{
		ID:          p.ID,
		TimeCreated: p.TimeCreated.Format(time.RFC3339),
		Title:       p.Title,
		Questions:   questions,
	}
}

func NewPolls(ps []*model.Poll) []*Poll {
	polls := make([]*Poll, len(ps))

	for _, p := range ps {
		polls = append(polls, NewPoll(p))
	}

	return polls
}

func (p *Poll) ToModel() (*model.Poll, error) {
	ts, err := time.Parse(time.RFC3339, p.TimeCreated)
	if err != nil {
		return nil, fmt.Errorf("parsing ts: %w", err)
	}

	questions := make([]*model.Question, len(p.Questions))
	for _, q := range p.Questions {
		questions = append(questions, q.ToModel())
	}

	return &model.Poll{
		ID:          p.ID,
		TimeCreated: ts,
		Title:       p.Title,
		Questions:   questions,
	}, nil
}
