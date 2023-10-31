package dto

import (
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
