package handler

import (
	"encoding/json"
	"errors"
	"github.com/MartinRusso28/go-error_handling/cmd/app/internal/animal"
	"github.com/MartinRusso28/go-error_handling/cmd/app/internal/animal/database"
	"io/ioutil"
	"net/http"
)

type Body struct {
	Code string `json:"code"`
	User int `json:"user"`
}

func SearchAnimal(w http.ResponseWriter, r *http.Request) error {
	var (
		bytes []byte
		b     Body
	)

	ctx := r.Context()

	db := database.Database{}

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	if err := json.Unmarshal(bytes, &b); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	a, err := animal.SearchAnimalByCode(ctx, &db, b.User, b.Code)

	if err != nil{
		if errors.Is(err, animal.ErrUnauthorized) {
			w.WriteHeader(http.StatusUnauthorized)
			return err
		}

		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(a))

	return nil
}