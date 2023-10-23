package questions

import "github.com/horockey/poller_backend/internal/model"

type Repository interface {
	Get(id string) (*model.Question, error)
	Add(q *model.Question) (*model.Question, error)
	Edit(id string, q *model.Question) (*model.Question, error)
	Delete(id string) (*model.Question, error)
}
