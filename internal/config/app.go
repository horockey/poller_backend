package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Address             string
	ShutdownTimeoutMsec int

	HashSeed string
}

func New() (*Config, error) {
	// TODO: refactor
	cfg := Config{}

	cfg.Address = os.Getenv("POLLER_BACKEND_ADDRESS")
	if cfg.Address == "" {
		return nil, errors.New("missing POLLER_BACKEND_ADDRESS env")
	}

	cfg.HashSeed = os.Getenv("POLLER_BACKEND_HASHSEED")
	if cfg.HashSeed == "" {
		return nil, errors.New("missing POLLER_BACKEND_HASHSEED env")
	}

	timeoutStr := os.Getenv("POLLER_BACKEND_SHUTDOWN_TIMEOUT_MSEC")
	if timeoutStr == "" {
		return nil, errors.New("missing POLLER_BACKEND_SHUTDOWN_TIMEOUT_MSEC env")
	}

	var err error
	cfg.ShutdownTimeoutMsec, err = strconv.Atoi(timeoutStr)
	if err != nil {
		return nil, fmt.Errorf("casting timeout to int: %w", err)
	}

	return &cfg, nil
}
