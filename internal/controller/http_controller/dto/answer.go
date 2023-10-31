package dto

import "github.com/horockey/poller_backend/internal/model"

type Answer struct {
	Text string `json:"text"`
}

func NewAnswer(a *model.Answer) *Answer {
	return &Answer{
		Text: a.Text,
	}
}
