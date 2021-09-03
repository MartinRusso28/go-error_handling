package test

import (
	"github.com/MartinRusso28/go-error_handling/cmd/app/internal/handler"
	"net/http"
	"net/http/httptest"
)

func SearchAnimalTest() {
	server := httptest.NewServer(http.HandlerFunc(handler.SearchAnimal))

	server.Start()

	defer server.Close()
}
//
//func doLocalRequest(method string, url string, headers http.Header, body string) *testResponse {
//	var reader io.Reader = nil
//
//	if body != "" {
//		reader = strings.NewReader(body)
//	}
//
//	request, _ := http.NewRequest(method, url, reader)
//
//	request.Header = headers
//
//	response := httptest.NewRecorder()
//
//	app.ServeHTTP(response, request)
//
//	var b []byte
//	if response.Body != nil {
//		b = response.Body.Bytes()
//	} else {
//		b = []byte("API did not respond to the call") //Not the same as an empty response
//	}
//
//	return &testResponse{HttpStatusCode: response.Code, Headers: response.Header(), Body: b}
//}
