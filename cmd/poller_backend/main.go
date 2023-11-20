package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/horockey/poller_backend/internal/adapter/event/attempt_action/zerolog_attempt_action"
	"github.com/horockey/poller_backend/internal/adapter/event/poll_action/zerolog_poll_action"
	"github.com/horockey/poller_backend/internal/adapter/repository/attempts/inmemory_attempts"
	"github.com/horockey/poller_backend/internal/adapter/repository/polls/inmemory_polls"
	"github.com/horockey/poller_backend/internal/config"
	"github.com/horockey/poller_backend/internal/controller/http_controller"
	attempts_management_impl "github.com/horockey/poller_backend/internal/usecase/attempts_management/impl"
	polls_management_impl "github.com/horockey/poller_backend/internal/usecase/polls_management/impl"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}).With().Timestamp().Logger()

	cfg, err := config.New()
	if err != nil {
		logger.Fatal().
			Err(fmt.Errorf("creating config: %w", err)).
			Send()
	}

	// Adapters
	attemptActionEvent := zerolog_attempt_action.New(logger)
	pollActionEvent := zerolog_poll_action.New(logger)

	attemptsRepo := inmemory_attempts.New()
	pollsRepo := inmemory_polls.New()

	// Usecase
	attemptsUC := attempts_management_impl.New(attemptsRepo, attemptActionEvent, logger)
	pollsUC := polls_management_impl.New(pollsRepo, pollActionEvent, logger)

	// Controller
	ctrl := http_controller.New(
		cfg.Address,
		time.Duration(cfg.ShutdownTimeoutMsec)*time.Millisecond,
		[]byte(cfg.HashSeed),
		attemptsUC,
		pollsUC,
		logger,
	)

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGKILL,
	)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		ctrl.Start(ctx)
	}()

	logger.Info().Msg("service started")
	wg.Wait()
	logger.Info().Msg("service stopped")
}
