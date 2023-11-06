package inmemory_attempts

import (
	"sync"

	"github.com/horockey/poller_backend/internal/adapter/repository"
	"github.com/horockey/poller_backend/internal/adapter/repository/attempts"
	"github.com/horockey/poller_backend/internal/model"
)

var _ attempts.Repository = &inmemoryAttempts{}

type inmemoryAttempts struct {
	mu sync.RWMutex

	storage map[string][]*model.Attempt
}

func New() *inmemoryAttempts {
	return &inmemoryAttempts{
		storage: map[string][]*model.Attempt{},
	}
}

func (repo *inmemoryAttempts) Get(id string) (*model.Attempt, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	for _, ats := range repo.storage {
		for _, at := range ats {
			if at.ID == id {
				return at, nil
			}
		}
	}

	return nil, repository.ErrNotFound
}

func (repo *inmemoryAttempts) GetAllForPoll(pollId string) ([]*model.Attempt, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	ats, found := repo.storage[pollId]
	if !found {
		return nil, repository.ErrNotFound
	}

	return ats, nil
}

func (repo *inmemoryAttempts) Add(a *model.Attempt) (*model.Attempt, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	attempts.FixID(a)

	for _, at := range repo.storage[a.PollID] {
		if at.ID == a.ID {
			return nil, repository.ErrAlreadyExists
		}
	}

	repo.storage[a.PollID] = append(repo.storage[a.PollID], a)

	return a, nil
}

func (repo *inmemoryAttempts) Delete(id string) (*model.Attempt, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for _, ats := range repo.storage {
		for idx, at := range ats {
			if at.ID == id {
				if idx < len(ats)-1 {
					ats = append(ats[:idx], ats[idx+1:]...)
				} else {
					ats = ats[:idx]
				}

				return at, nil
			}
		}
	}

	return nil, repository.ErrNotFound
}

func (repo *inmemoryAttempts) DeleteAllForPoll(pollId string) ([]*model.Attempt, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	ats, found := repo.storage[pollId]
	if !found {
		return nil, repository.ErrNotFound
	}
	res := make([]*model.Attempt, len(ats))

	copy(res, ats)
	delete(repo.storage, pollId)

	return res, nil
}
