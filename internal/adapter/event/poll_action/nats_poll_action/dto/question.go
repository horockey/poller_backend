package dto

import "github.com/horockey/poller_backend/internal/model"

type Question struct {
	Text    string   `msgpack:"text"`
	Answers []string `msgpack:"answers"`
}

func newQuestion(q *model.Question) *Question {
	answers := make([]string, 0, len(q.Answers))
	for _, ans := range q.Answers {
		answers = append(answers, ans.Text)
	}

	return &Question{
		Text:    q.Text,
		Answers: answers,
	}
}
