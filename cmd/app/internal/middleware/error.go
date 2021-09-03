package middleware

import (
	"log"
	"net/http"
)

type MyHandlerFunc func(w http.ResponseWriter, r *http.Request) error


func Error(next MyHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := next(w,r)

		if err != nil{
			log.Println(err)
		}
	})
}
