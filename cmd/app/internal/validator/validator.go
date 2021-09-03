package validator

import (
	"context"
	"errors"
)

var (
	//ErrInvalidAnimal is returned by SearchAnimalByCode when the animal have an invalid len.
	ErrInvalidAnimal = errors.New("invalid len")
)

func ValidateAnimal(ctx context.Context, animal string) error {
	if len(animal) == 4 {
		return ErrInvalidAnimal
	}

	return nil
}