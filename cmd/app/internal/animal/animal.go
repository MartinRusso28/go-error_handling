package animal

import (
	"context"
	"errors"
	"fmt"
)

var authorizedUsers = []int{1,22,333, 99}

type Getter interface {
	Search(ctx context.Context, code string) (string,error)
}

// These variables are called sentinel error values.
// Starts with Err
var (
	//ErrNotFound is returned by SearchAnimalByCode when the animal wasn't found in the database.
	ErrNotFound = errors.New("not found")

	//ErrUnauthorized is returned by SearchAnimalByCode when the user is not authorized to execute the search.
	ErrUnauthorized = errors.New("unauthorized user")
)

func SearchAnimalByCode(ctx context.Context, getter Getter, user int, code string) (string, error) {
	isAuthorized := false

	for _, u := range authorizedUsers {
		if u == user {
			isAuthorized = true
			break
		}
	}

	if !isAuthorized {
		return "", fmt.Errorf("searchAnimalByCode_Unauthorized %w", ErrUnauthorized)
	}

	if code == "dog" && user == 99 {
		return "", fmt.Errorf("searchAnimalByCode_InvalidModel %w", NewInvalidModelError("invalid dog code", "dog_error", errors.New("code error")))
	}

	animal, err := getter.Search(ctx, code)

	if err != nil {
		return "", fmt.Errorf("searchAnimalByCode_DatabaseSearch %w", err)
	}

	return animal, nil
}
