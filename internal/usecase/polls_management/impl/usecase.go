package polls_management_impl

import (
	"fmt"

	"github.com/horockey/poller_backend/internal/adapter/event"
	"github.com/horockey/poller_backend/internal/adapter/event/poll_action"
	"github.com/horockey/poller_backend/internal/adapter/repository/polls"
	"github.com/horockey/poller_backend/internal/model"
	"github.com/horockey/poller_backend/internal/usecase/polls_management"
	"github.com/rs/zerolog"
)

var _ polls_management.Usecase = &pollsManagement{}

type pollsManagement struct {
	repo polls.Repository
	ev   poll_action.Event

	logger zerolog.Logger
}

func New(
	repo polls.Repository,
	ev poll_action.Event,
	logger zerolog.Logger,
) *pollsManagement {
	return &pollsManagement{
		repo:   repo,
		ev:     ev,
		logger: logger,
	}
}

func (uc *pollsManagement) Get(id string) (*model.Poll, error) {
	p, err := uc.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("getting poll from repo: %w", err)
	}

	if err := uc.ev.Send(p, event.ActionRead); err != nil {
		uc.logger.Error().
			Err(fmt.Errorf("sending event: %w", err)).
			Send()
	}

	return p, nil
}

func (uc *pollsManagement) GetAll() ([]*model.Poll, error) {
	ps, err := uc.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("getting poll from repo: %w", err)
	}

	for _, p := range ps {
		if err := uc.ev.Send(p, event.ActionRead); err != nil {
			uc.logger.Error().
				Err(fmt.Errorf("sending event: %w", err)).
				Send()
		}
	}

	return ps, nil
}

func (uc *pollsManagement) Add(poll *model.Poll) (*model.Poll, error) {
	p, err := uc.repo.Add(poll)
	if err != nil {
		return nil, fmt.Errorf("getting poll from repo: %w", err)
	}

	if err := uc.ev.Send(p, event.ActionRead); err != nil {
		uc.logger.Error().
			Err(fmt.Errorf("sending event: %w", err)).
			Send()
	}

	return p, nil
}

func (uc *pollsManagement) Delete(id string) (*model.Poll, error) {
	p, err := uc.repo.Delete(id)
	if err != nil {
		return nil, fmt.Errorf("getting poll from repo: %w", err)
	}

	if err := uc.ev.Send(p, event.ActionRead); err != nil {
		uc.logger.Error().
			Err(fmt.Errorf("sending event: %w", err)).
			Send()
	}

	return p, nil
}
