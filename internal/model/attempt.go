package model

import "time"

type Attempt struct {
	ID      string
	Ts      time.Time
	PollID  string
	Answers []*AttemptAnswer
}

type AttemptAnswer struct {
	Question *Question
	Answers  []*Answer
}
