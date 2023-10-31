package model

type Poll struct {
	ID        string
	Title     string
	Questions []*Question
}
