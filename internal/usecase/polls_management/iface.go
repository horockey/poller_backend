package polls_management

import "github.com/horockey/poller_backend/internal/model"

type Usecase interface {
	Get(id string) (*model.Poll, error)
	GetAll() ([]*model.Poll, error)
	Add(p *model.Poll) (*model.Poll, error)
	Delete(id string) (*model.Poll, error)
}
