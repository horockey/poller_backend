package inmemory_polls

import (
	"sync"

	"github.com/horockey/poller_backend/internal/adapter/repository"
	"github.com/horockey/poller_backend/internal/adapter/repository/polls"
	"github.com/horockey/poller_backend/internal/model"
)

var _ polls.Repository = &inmemoryPolls{}

type inmemoryPolls struct {
	mu sync.RWMutex

	storage map[string]*model.Poll
}

func New() *inmemoryPolls {
	return &inmemoryPolls{
		storage: map[string]*model.Poll{},
	}
}

func (repo *inmemoryPolls) Get(id string) (*model.Poll, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	poll, found := repo.storage[id]
	if !found {
		return nil, repository.ErrNotFound
	}

	return poll, nil
}

func (repo *inmemoryPolls) GetAll() ([]*model.Poll, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	res := make([]*model.Poll, 0, len(repo.storage))
	for _, poll := range repo.storage {
		res = append(res, poll)
	}

	return res, nil
}

func (repo *inmemoryPolls) Add(p *model.Poll) (*model.Poll, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	polls.FixID(p)

	if _, found := repo.storage[p.ID]; found {
		return nil, repository.ErrAlreadyExists
	}

	repo.storage[p.ID] = p

	return nil, nil
}

func (repo *inmemoryPolls) Delete(id string) (*model.Poll, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	poll, found := repo.storage[id]
	if !found {
		return nil, repository.ErrNotFound
	}

	delete(repo.storage, id)

	return poll, nil
}
