package model

type Question struct {
	ID         string
	Text       string
	AnswersIDs []string
}
