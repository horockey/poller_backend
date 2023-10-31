package http_controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horockey/go-toolbox/http_helpers"
	"github.com/horockey/poller_backend/internal/controller/http_controller/dto"
)

func (ctrl *httpController) attemptPollIdDelete(w http.ResponseWriter, req *http.Request) {
	if err := ctrl.checkAuthToken(req); err != nil {
		resErr := fmt.Errorf("checking auth token: %w", err)

		if err := http_helpers.RespondWithErr(
			w,
			http.StatusForbidden,
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

func (ctrl *httpController) attemptPollIdGet(w http.ResponseWriter, req *http.Request) {
	if err := ctrl.checkAuthToken(req); err != nil {
		resErr := fmt.Errorf("checking auth token: %w", err)

		if err := http_helpers.RespondWithErr(
			w,
			http.StatusForbidden,
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

func (ctrl *httpController) attemptPollIdPost(w http.ResponseWriter, req *http.Request) {
	if err := ctrl.checkAuthToken(req); err != nil {
		resErr := fmt.Errorf("checking auth token: %w", err)

		if err := http_helpers.RespondWithErr(
			w,
			http.StatusForbidden,
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

	atReq := dto.Attempt{}
	if err := json.NewDecoder(req.Body).Decode(&atReq); err != nil {
		resErr := fmt.Errorf("decoding json from req body: %w", err)

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

	if atReq.PollID != pollID {
		resErr := fmt.Errorf(
			"pollId in body (%s) does not match PollId in param (%s)",
			atReq.PollID,
			pollID,
		)

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

	at, err := atReq.ToModel()
	if err != nil {
		resErr := fmt.Errorf("converting attempt to model: %w", err)

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

	resAt, err := ctrl.attempts.Add(at)
	if err != nil {
		resErr := fmt.Errorf("calling usecase: %w", err)

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

	if err := http_helpers.RespondOK(w, dto.NewAttempt(resAt)); err != nil {
		ctrl.logger.Error().
			Err(fmt.Errorf("responding: %w", err)).
			Send()
	}
}

func (ctrl *httpController) attemptPollIdAttemptIdDelete(w http.ResponseWriter, req *http.Request) {
	if err := ctrl.checkAuthToken(req); err != nil {
		resErr := fmt.Errorf("checking auth token: %w", err)

		if err := http_helpers.RespondWithErr(
			w,
			http.StatusForbidden,
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

func (ctrl *httpController) attemptPollIdAttemptIdGet(w http.ResponseWriter, req *http.Request) {
	if err := ctrl.checkAuthToken(req); err != nil {
		resErr := fmt.Errorf("checking auth token: %w", err)

		if err := http_helpers.RespondWithErr(
			w,
			http.StatusForbidden,
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
