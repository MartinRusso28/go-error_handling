package main

import (
	"fmt"
	"github.com/MartinRusso28/go-error_handling/cmd/app/internal/handler"
	"github.com/MartinRusso28/go-error_handling/cmd/app/internal/middleware"
	"log"
	"net/http"
)

func main() {
	server := http.NewServeMux()

	searchAnimalHandler := handler.SearchAnimal

	server.Handle("/animal", middleware.Error(searchAnimalHandler))

	fmt.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal(err)
	}
}
