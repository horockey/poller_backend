package model

type Attempt struct {
	ID         string
	PollID     string
	AnswersIDs []string
}
