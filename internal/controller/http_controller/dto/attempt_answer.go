package dto

import "github.com/horockey/poller_backend/internal/model"

type AttemptAnswer struct {
	Question *Question `json:"question"`
	Answers  []*Answer `json:"answers"`
}

func newAttemptAnswer(a *model.AttemptAnswer) *AttemptAnswer {
	answers := make([]*Answer, len(a.Answers))
	for _, ans := range a.Answers {
		answers = append(answers, newAnswer(ans))
	}

	return &AttemptAnswer{
		Question: newQuestion(a.Question),
		Answers:  answers,
	}
}

func (a *AttemptAnswer) ToModel() *model.AttemptAnswer {
	answers := make([]*model.Answer, len(a.Answers))
	for _, ans := range a.Answers {
		answers = append(answers, ans.ToModel())
	}

	return &model.AttemptAnswer{
		Question: a.Question.ToModel(),
		Answers:  answers,
	}
}
