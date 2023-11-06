package inmemory_polls

import (
	"github.com/horockey/poller_backend/internal/adapter/repository"
	"github.com/horockey/poller_backend/internal/adapter/repository/polls"
	"github.com/horockey/poller_backend/internal/model"
)

var _ polls.Repository = &inmemoryPolls{}

type inmemoryPolls struct {
	storage map[string]*model.Poll
}

func New() *inmemoryPolls {
	return &inmemoryPolls{
		storage: map[string]*model.Poll{},
	}
}

func (repo *inmemoryPolls) Get(id string) (*model.Poll, error) {
	poll, found := repo.storage[id]
	if !found {
		return nil, repository.ErrNotFound
	}

	return poll, nil
}

func (repo *inmemoryPolls) GetAll() ([]*model.Poll, error) {
	// TODO
	return nil, nil
}

func (repo *inmemoryPolls) Add(p *model.Poll) (*model.Poll, error) {
	// TODO
	return nil, nil
}

func (repo *inmemoryPolls) Delete(id string) (*model.Poll, error) {
	// TODO
	return nil, nil
}
