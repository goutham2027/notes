package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

type test struct {
	url                string
	expectedStatusCode int
}

func TestService(t *testing.T) {
	server := NewServer()
	router := mux.NewRouter()
	server.Register(router)
	var tests []test
	tests = append(tests, test{url: "/", expectedStatusCode: http.StatusOK})
	tests = append(tests, test{url: "/abc", expectedStatusCode: http.StatusNotFound})
	for _, test := range tests {
		req := httptest.NewRequest("GET", test.url, strings.NewReader("empty body"))
		rw := httptest.NewRecorder()
		router.ServeHTTP(rw, req)
		response := rw.Result()
		statusCode := response.StatusCode

		if statusCode != test.expectedStatusCode {
			t.Errorf("Status code didn't match, expected: %d, got: %d", test.expectedStatusCode, statusCode)
		}

		// body, err := io.ReadAll(response.Body)
		// if err != nil {
		// 	fmt.Println("Error reading response body", err)
		// 	return
		// }
		// bodyString := string(body)
		// if bodyString != "hello world!" {
		// 	t.Errorf("expected: %q, got: %q", "hello world!", bodyString)
		// }
	}

}
