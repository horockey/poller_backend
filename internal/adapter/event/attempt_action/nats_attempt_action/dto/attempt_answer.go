package dto

import "github.com/horockey/poller_backend/internal/model"

type AttemptAnswer struct {
	Question string   `msgpack:"question"`
	Answers  []string `msgpack:"answers"`
}

func newAttemptAnswer(a *model.AttemptAnswer) *AttemptAnswer {
	answers := make([]string, 0, len(a.Answers))
	for _, ans := range a.Answers {
		answers = append(answers, ans.Text)
	}

	return &AttemptAnswer{
		Question: a.Question.Text,
		Answers:  answers,
	}
}
