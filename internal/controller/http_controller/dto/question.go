package dto

import "github.com/horockey/poller_backend/internal/model"

type Question struct {
	Text    string    `json:"question"`
	Answers []*Answer `json:"answers"`
}

func newQuestion(q *model.Question) *Question {
	answers := make([]*Answer, len(q.Answers))
	for _, ans := range q.Answers {
		answers = append(answers, newAnswer(ans))
	}

	return &Question{
		Text:    q.Text,
		Answers: answers,
	}
}

func (q *Question) ToModel() *model.Question {
	answers := make([]*model.Answer, len(q.Answers))
	for _, ans := range q.Answers {
		answers = append(answers, ans.ToModel())
	}

	return &model.Question{
		Text:    q.Text,
		Answers: answers,
	}
}
