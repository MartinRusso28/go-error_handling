package handler

import "fmt"

// type error values
// por convención, terminan con Error.

type InvalidModelError struct {
	name string
	errorCode int
	err error
}

func NewInvalidModelError(name string, code int, err error) InvalidModelError{
	return InvalidModelError{
		name: name,
		errorCode: code,
		err: err,
	}
}

func (e *InvalidModelError) Error() string {
	if e.errorCode == 0 {
		return fmt.Sprintf("invalid model error code")
	}

	return fmt.Sprintf("the value %s with code %d is invalid", e.name, e.errorCode)
}

func (e *InvalidModelError) Unwrap() error {
	return e.err
}