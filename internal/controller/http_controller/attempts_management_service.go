package http_controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/horockey/go-toolbox/http_helpers"
	"github.com/horockey/poller_backend/internal/controller/http_controller/dto"
)

func (ctrl *httpController) PruneAttemptsDelete(w http.ResponseWriter, req *http.Request) {
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

	ats, err := ctrl.attempts.Prune()
	if err != nil {
		resErr := fmt.Errorf("pruning attempts: %w", err)

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
