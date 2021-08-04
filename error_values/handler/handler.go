package handler

import "errors"

// sentinel error values.
// por convención siempre empiezan con Err
var (
	ErrNotFound = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalidHandler = errors.New("invalid handler for user")
)

func HandleUserAndModel(user int, model string) error {

}

