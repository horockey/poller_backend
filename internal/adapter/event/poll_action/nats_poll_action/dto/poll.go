package dto

import (
	"time"

	"github.com/horockey/poller_backend/internal/model"
)

type Poll struct {
	ID          string      `msgpack:"id"`
	TimeCreated time.Time   `msgpack:"time_created"`
	Title       string      `msgpack:"title"`
	Questions   []*Question `msgpack:"questions"`
}

func newPoll(p *model.Poll) *Poll {
	questions := make([]*Question, 0, len(p.Questions))
	for _, q := range p.Questions {
		questions = append(questions, newQuestion(q))
	}

	return &Poll{
		ID:          p.ID,
		TimeCreated: p.TimeCreated,
		Title:       p.Title,
		Questions:   questions,
	}
}
