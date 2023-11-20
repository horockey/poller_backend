package http_controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/horockey/poller_backend/internal/usecase/attempts_management"
	"github.com/horockey/poller_backend/internal/usecase/polls_management"
	"github.com/rs/zerolog"
)

type httpController struct {
	serv            *http.Server
	shutdownTimeout time.Duration

	hashSeed []byte

	attempts attempts_management.Usecase
	polls    polls_management.Usecase

	logger zerolog.Logger
}

func New(
	addr string,
	shutdownTimeout time.Duration,
	hashSeed []byte,
	attempts attempts_management.Usecase,
	polls polls_management.Usecase,
	logger zerolog.Logger,
) *httpController {
	ctrl := httpController{
		serv: &http.Server{
			Addr: addr,
		},
		shutdownTimeout: shutdownTimeout,
		hashSeed:        hashSeed,
		attempts:        attempts,
		polls:           polls,
		logger:          logger,
	}
	ctrl.serv.Handler = ctrl.newRouter()

	return &ctrl
}

func (ctrl *httpController) Start(ctx context.Context) error {
	ctrl.logger.Info().Str("addr", ctrl.serv.Addr).Msg("HTTP controller started")

	errs := make(chan error)
	go func() {
		err := ctrl.serv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errs <- err
		}
	}()

	select {
	case err := <-errs:
		return fmt.Errorf("running http server: %w", err)
	case <-ctx.Done():
		var resErr error
		if !errors.Is(ctx.Err(), context.Canceled) {
			resErr = fmt.Errorf("running context: %w", resErr)
		}

		sdCtx, cancel := context.WithTimeout(context.TODO(), ctrl.shutdownTimeout)
		defer cancel()

		if err := ctrl.serv.Shutdown(sdCtx); err != nil {
			resErr = errors.Join(resErr, fmt.Errorf("shutting down server: %w", err))
		}
		return resErr
	}
}
