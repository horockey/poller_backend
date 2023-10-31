package attempts

import "github.com/horockey/poller_backend/internal/model"

type Repository interface {
	Get(id string) (*model.Attempt, error)
	GetAllForPoll(pollId string) ([]*model.Attempt, error)
	Add(a *model.Attempt) (*model.Attempt, error)
	Delete(id string) (*model.Attempt, error)
	DeleteAllForPoll(pollId string) ([]*model.Attempt, error)
	Prune() ([]*model.Attempt, error)
}
