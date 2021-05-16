package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var s *server

func TestMain(m *testing.M) {
	s, _ = new(":7000")
	code := m.Run()
	os.Exit(code)
}

// executeRequest takes a request and records the response
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.Get("/{statusCode}", statusCodeHandler(s))
	s.Router.ServeHTTP(rr, req)
	return rr
}

func TestReturnInvalidCode(t *testing.T) {
	t.Parallel()

	req, _ := http.NewRequest("GET", "/abc", nil)
	res := executeRequest(req)

	assert.Equal(t, http.StatusBadRequest, res.Code, "should return a 400 status")
	assert.Equal(t, "Invalid status code\n", res.Body.String(), "should return invalid status code")
}

func TestReturnGoCode(t *testing.T) {
	t.Parallel()

	req, _ := http.NewRequest("GET", "/404", nil)
	res := executeRequest(req)

	assert.Equal(t, http.StatusNotFound, res.Code, "should return a 404 status")
	assert.Equal(t, "Not Found\n", res.Body.String(), "should return Not Found body")
}

func TestReturnCustomCode(t *testing.T) {
	t.Parallel()

	req, _ := http.NewRequest("GET", "/218", nil)
	res := executeRequest(req)

	assert.Equal(t, 218, res.Code, "should return a 404 status")
	assert.Equal(t, "This is fine\n", res.Body.String(), "should return this is fine body")
}

func TestReturnUnknownCode(t *testing.T) {
	t.Parallel()

	req, _ := http.NewRequest("GET", "/999", nil)
	res := executeRequest(req)

	assert.Equal(t, 999, res.Code, "should return a unknown status")
	assert.Equal(t, "\n", res.Body.String(), "should return this is fine body")
}
