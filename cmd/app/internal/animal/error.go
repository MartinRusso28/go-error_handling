package animal

import "fmt"

// InvalidModelError these types are called error type values.
// Ends wit "Error".
type InvalidModelError struct {
	name string
	errorCode string
	err error
}

func NewInvalidModelError(name string, code string, err error) *InvalidModelError {
	return &InvalidModelError{
		name: name,
		errorCode: code,
		err: err,
	}
}

func (e *InvalidModelError) Error() string {
	if e.errorCode == "0" {
		return fmt.Sprintf("invalid model example code")
	}

	return fmt.Sprintf("the value %s with code %s is invalid", e.name, e.errorCode)
}

func (e *InvalidModelError) Unwrap() error {
	return e.err
}