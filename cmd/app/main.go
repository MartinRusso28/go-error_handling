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

	server.Handle("/animal", middleware.Error(handler.SearchAnimal))

	fmt.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal(err)
	}
}
