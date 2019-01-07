package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayer(t *testing.T) {
	t.Run("return Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest()
		response := httptest.NewRecorder()

		GetScore(response, request)

		assertResponse(t, response.Body.String(), "20")
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest()
		response := httptest.NewRecorder()

		GetScore(response, request)

		assertResponse(t, response.Body.String(), "10")
	})
}

func newGetScoreRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
	return req
}

func assertResponse(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
