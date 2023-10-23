package polls

import "github.com/horockey/poller_backend/internal/model"

type Repository interface {
	Get(id string) (*model.Poll, error)
	Add(p *model.Poll) (*model.Poll, error)
	Edit(id string, p *model.Poll) (*model.Poll, error)
	Delete(id string) (*model.Poll, error)
}
