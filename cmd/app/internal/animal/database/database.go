package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/MartinRusso28/go-error_handling/cmd/app/internal/validator"
)

var database = map[int]string{1:"dog", 22:"cat", 33:"dogg"}

var (
	ErrNotFound = errors.New("animal not found")
)

type Database struct{}

func (db *Database) Search(ctx context.Context, code string) (string,error) {
	animal := ""

	for _, v := range database {
		if code == v {
			animal = v
			break
		}
	}

	if animal == "" {
		return "", ErrNotFound
	}

	if err := validator.ValidateAnimal(ctx, animal); err != nil {
		return "", fmt.Errorf("search_ValidateAnimal %w", err)
	}

	return animal, nil
}
