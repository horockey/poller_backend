package attempts_management_impl

import (
	"fmt"

	"github.com/horockey/poller_backend/internal/adapter/event"
	"github.com/horockey/poller_backend/internal/adapter/event/attempt_action"
	"github.com/horockey/poller_backend/internal/adapter/repository/attempts"
	"github.com/horockey/poller_backend/internal/model"
	"github.com/horockey/poller_backend/internal/usecase/attempts_management"
	"github.com/rs/zerolog"
)

var _ attempts_management.Usecase = &attemptsManagement{}

type attemptsManagement struct {
	repo attempts.Repository
	ev   attempt_action.Event

	logger zerolog.Logger
}

func New(
	repo attempts.Repository,
	ev attempt_action.Event,
	logger zerolog.Logger,
) *attemptsManagement {
	return &attemptsManagement{
		repo:   repo,
		ev:     ev,
		logger: logger,
	}
}

func (uc *attemptsManagement) Get(id string) (*model.Attempt, error) {
	at, err := uc.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("getting attempt from repo: %w", err)
	}

	if err := uc.ev.Send(at, event.ActionRead); err != nil {
		uc.logger.Error().
			Err(fmt.Errorf("sending event: %w", err)).
			Send()
	}

	return at, nil
}

func (uc *attemptsManagement) GetAllForPoll(pollId string) ([]*model.Attempt, error) {
	ats, err := uc.repo.GetAllForPoll(pollId)
	if err != nil {
		return nil, fmt.Errorf("getting attempts from repo: %w", err)
	}

	for _, at := range ats {
		if err := uc.ev.Send(at, event.ActionRead); err != nil {
			uc.logger.Error().
				Err(fmt.Errorf("sending event: %w", err)).
				Send()
		}
	}

	return ats, nil
}

func (uc *attemptsManagement) Add(a *model.Attempt) (*model.Attempt, error) {
	at, err := uc.repo.Add(a)
	if err != nil {
		return nil, fmt.Errorf("adding attempt to repo: %w", err)
	}

	if err := uc.ev.Send(at, event.ActionCreate); err != nil {
		uc.logger.Error().
			Err(fmt.Errorf("sending event: %w", err)).
			Send()
	}

	return at, nil
}

func (uc *attemptsManagement) Delete(id string) (*model.Attempt, error) {
	at, err := uc.repo.Delete(id)
	if err != nil {
		return nil, fmt.Errorf("deleting attempt from repo: %w", err)
	}

	if err := uc.ev.Send(at, event.ActionDelete); err != nil {
		uc.logger.Error().
			Err(fmt.Errorf("sending event: %w", err)).
			Send()
	}

	return at, nil
}

func (uc *attemptsManagement) DeleteAllForPoll(pollId string) ([]*model.Attempt, error) {
	ats, err := uc.repo.DeleteAllForPoll(pollId)
	if err != nil {
		return nil, fmt.Errorf("deleting attempts from repo: %w", err)
	}

	for _, at := range ats {
		if err := uc.ev.Send(at, event.ActionDelete); err != nil {
			uc.logger.Error().
				Err(fmt.Errorf("sending event: %w", err)).
				Send()
		}
	}

	return ats, nil
}

func (uc *attemptsManagement) Prune() ([]*model.Attempt, error) {
	ats, err := uc.repo.Prune()
	if err != nil {
		return nil, fmt.Errorf("pruning attempts from repo: %w", err)
	}

	for _, at := range ats {
		if err := uc.ev.Send(at, event.ActionDelete); err != nil {
			uc.logger.Error().
				Err(fmt.Errorf("sending event: %w", err)).
				Send()
		}
	}

	return ats, nil
}
