package model

import "time"

type Attempt struct {
	ID      string
	Ts      time.Time
	PollID  string
	Answers [][]*Answer
}
