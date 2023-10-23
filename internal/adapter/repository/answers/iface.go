package answers

import "github.com/horockey/poller_backend/internal/model"

type Repository interface {
	Get(id string) (*model.Answer, error)
	Add(a *model.Answer) (*model.Answer, error)
	Edit(id string, a *model.Answer) (*model.Answer, error)
	Delete(id string) (*model.Answer, error)
}
