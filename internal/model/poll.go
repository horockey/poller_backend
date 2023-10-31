package model

import "time"

type Poll struct {
	ID          string
	TimeCreated time.Time
	Title       string
	Questions   []*Question
}
