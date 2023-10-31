package dto

import "github.com/horockey/poller_backend/internal/model"

type Question struct {
	Text    string    `json:"question"`
	Answers []*Answer `json:"answers"`
}

func NewQuestion(q *model.Question) *Question {
	answers := make([]*Answer, len(q.Answers))
	for _, ans := range q.Answers {
		answers = append(answers, NewAnswer(ans))
	}

	return &Question{
		Text:    q.Text,
		Answers: answers,
	}
}
