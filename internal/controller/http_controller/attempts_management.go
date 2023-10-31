package http_controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horockey/go-toolbox/http_helpers"
	"github.com/horockey/poller_backend/internal/controller/http_controller/dto"
)

func (ctrl *httpController) AttemptPollIdDelete(w http.ResponseWriter, req *http.Request) {
	if err := ctrl.checkAuthToken(req); err != nil {
		resErr := fmt.Errorf("checking auth token: %w", err)

		if err := http_helpers.RespondWithErr(
			w,
			http.StatusBadRequest,
			resErr,
		); err != nil {
			resErr = errors.Join(resErr, fmt.Errorf("responing: %w", err))
		}

		ctrl.logger.Error().
			Err(resErr).
			Send()
		return
	}

	pollID, found := mux.Vars(req)[pollIDParamName]
	if !found {
		resErr := errors.New("poll id is empty")

		if err := http_helpers.RespondWithErr(
			w,
			http.StatusBadRequest,
			resErr,
		); err != nil {
			resErr = errors.Join(resErr, fmt.Errorf("responing: %w", err))
		}

		ctrl.logger.Error().
			Err(resErr).
			Send()
	}

	ats, err := ctrl.attempts.DeleteAllForPoll(pollID)
	if err != nil {
		resErr := fmt.Errorf("calling usecase: %w", err)

		if err := http_helpers.RespondWithErr(
			w,
			http.StatusInternalServerError,
			resErr,
		); err != nil {
			resErr = errors.Join(resErr, fmt.Errorf("responing: %w", err))
		}

		ctrl.logger.Error().
			Err(resErr).
			Send()
		return
	}

	if err := http_helpers.RespondOK(w, dto.NewAttempts(ats)); err != nil {
		ctrl.logger.Error().
			Err(fmt.Errorf("responding: %w", err)).
			Send()
	}
}

func (ctrl *httpController) AttemptPollIdGet(w http.ResponseWriter, req *http.Request) {
	if err := ctrl.checkAuthToken(req); err != nil {
		resErr := fmt.Errorf("checking auth token: %w", err)

		if err := http_helpers.RespondWithErr(
			w,
			http.StatusBadRequest,
			resErr,
		); err != nil {
			resErr = errors.Join(resErr, fmt.Errorf("responing: %w", err))
		}

		ctrl.logger.Error().
			Err(resErr).
			Send()
		return
	}

	pollID, found := mux.Vars(req)[pollIDParamName]
	if !found {
		resErr := errors.New("poll id is empty")

		if err := http_helpers.RespondWithErr(
			w,
			http.StatusBadRequest,
			resErr,
		); err != nil {
			resErr = errors.Join(resErr, fmt.Errorf("responing: %w", err))
		}

		ctrl.logger.Error().
			Err(resErr).
			Send()
	}

	ats, err := ctrl.attempts.GetAllForPoll(pollID)
	if err != nil {
		resErr := fmt.Errorf("calling usecase: %w", err)

		if err := http_helpers.RespondWithErr(
			w,
			http.StatusInternalServerError,
			resErr,
		); err != nil {
			resErr = errors.Join(resErr, fmt.Errorf("responing: %w", err))
		}

		ctrl.logger.Error().
			Err(resErr).
			Send()
		return
	}

	if err := http_helpers.RespondOK(w, dto.NewAttempts(ats)); err != nil {
		ctrl.logger.Error().
			Err(fmt.Errorf("responding: %w", err)).
			Send()
	}
}

func (ctrl *httpController) AttemptPollIdAttemptIdDelete(w http.ResponseWriter, req *http.Request) {
	if err := ctrl.checkAuthToken(req); err != nil {
		resErr := fmt.Errorf("checking auth token: %w", err)

		if err := http_helpers.RespondWithErr(
			w,
			http.StatusBadRequest,
			resErr,
		); err != nil {
			resErr = errors.Join(resErr, fmt.Errorf("responing: %w", err))
		}

		ctrl.logger.Error().
			Err(resErr).
			Send()
		return
	}

	attemptID, found := mux.Vars(req)[attemptIDParamName]
	if !found {
		resErr := errors.New("attempt id is empty")

		if err := http_helpers.RespondWithErr(
			w,
			http.StatusBadRequest,
			resErr,
		); err != nil {
			resErr = errors.Join(resErr, fmt.Errorf("responing: %w", err))
		}

		ctrl.logger.Error().
			Err(resErr).
			Send()
	}

	at, err := ctrl.attempts.Delete(attemptID)
	if err != nil {
		resErr := fmt.Errorf("calling usecase: %w", err)

		if err := http_helpers.RespondWithErr(
			w,
			http.StatusInternalServerError,
			resErr,
		); err != nil {
			resErr = errors.Join(resErr, fmt.Errorf("responing: %w", err))
		}

		ctrl.logger.Error().
			Err(resErr).
			Send()
		return
	}

	if err := http_helpers.RespondOK(w, dto.NewAttempt(at)); err != nil {
		ctrl.logger.Error().
			Err(fmt.Errorf("responding: %w", err)).
			Send()
	}
}

func (ctrl *httpController) AttemptPollIdAttemptIdGet(w http.ResponseWriter, req *http.Request) {
	if err := ctrl.checkAuthToken(req); err != nil {
		resErr := fmt.Errorf("checking auth token: %w", err)

		if err := http_helpers.RespondWithErr(
			w,
			http.StatusBadRequest,
			resErr,
		); err != nil {
			resErr = errors.Join(resErr, fmt.Errorf("responing: %w", err))
		}

		ctrl.logger.Error().
			Err(resErr).
			Send()
		return
	}

	attemptID, found := mux.Vars(req)[attemptIDParamName]
	if !found {
		resErr := errors.New("attempt id is empty")

		if err := http_helpers.RespondWithErr(
			w,
			http.StatusBadRequest,
			resErr,
		); err != nil {
			resErr = errors.Join(resErr, fmt.Errorf("responing: %w", err))
		}

		ctrl.logger.Error().
			Err(resErr).
			Send()
	}

	at, err := ctrl.attempts.Get(attemptID)
	if err != nil {
		resErr := fmt.Errorf("calling usecase: %w", err)

		if err := http_helpers.RespondWithErr(
			w,
			http.StatusInternalServerError,
			resErr,
		); err != nil {
			resErr = errors.Join(resErr, fmt.Errorf("responing: %w", err))
		}

		ctrl.logger.Error().
			Err(resErr).
			Send()
		return
	}

	if err := http_helpers.RespondOK(w, dto.NewAttempt(at)); err != nil {
		ctrl.logger.Error().
			Err(fmt.Errorf("responding: %w", err)).
			Send()
	}
}
