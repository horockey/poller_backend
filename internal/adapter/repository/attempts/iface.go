package attempts

import "github.com/horockey/poller_backend/internal/model"

type Repository interface {
	Get(id string) (*model.Attempt, error)
	Add(a *model.Attempt) (*model.Attempt, error)
	Edit(id string, a *model.Attempt) (*model.Attempt, error)
	Delete(id string) (*model.Attempt, error)
}
