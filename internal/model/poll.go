package model

import "time"

type Poll struct {
	ID           string
	Title        string
	QuestionsIDs []string
	Until        time.Time
}