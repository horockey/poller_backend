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

func (ctrl *httpController) pollGet(w http.ResponseWriter, req *http.Request) {
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

	ps, err := ctrl.polls.GetAll()
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

	if err := http_helpers.RespondOK(w, dto.NewPolls(ps)); err != nil {
		ctrl.logger.Error().
			Err(fmt.Errorf("responding: %w", err)).
			Send()
	}
}

func (ctrl *httpController) pollIdDelete(w http.ResponseWriter, req *http.Request) {
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

	p, err := ctrl.polls.Delete(pollID)
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

	if err := http_helpers.RespondOK(w, dto.NewPoll(p)); err != nil {
		ctrl.logger.Error().
			Err(fmt.Errorf("responding: %w", err)).
			Send()
	}
}

func (ctrl *httpController) pollIdGet(w http.ResponseWriter, req *http.Request) {
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

	p, err := ctrl.polls.Get(pollID)
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

	if err := http_helpers.RespondOK(w, dto.NewPoll(p)); err != nil {
		ctrl.logger.Error().
			Err(fmt.Errorf("responding: %w", err)).
			Send()
	}
}

func (ctrl *httpController) pollPost(w http.ResponseWriter, req *http.Request) {
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

	pollReq := dto.Poll{}
	if err := json.NewDecoder(req.Body).Decode(&pollReq); err != nil {
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

	poll, err := pollReq.ToModel()
	if err != nil {
		resErr := fmt.Errorf("converting poll to model: %w", err)

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

	resPoll, err := ctrl.polls.Add(poll)
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

	if err := http_helpers.RespondOK(w, dto.NewPoll(resPoll)); err != nil {
		ctrl.logger.Error().
			Err(fmt.Errorf("responding: %w", err)).
			Send()
	}
}
