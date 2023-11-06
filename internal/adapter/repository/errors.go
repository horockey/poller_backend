package repository

import "errors"

var (
	ErrNotFound      = errors.New("item not found")
	ErrAlreadyExists = errors.New("item already exists")
)
