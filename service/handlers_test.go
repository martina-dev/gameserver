package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayer(t *testing.T) {
	t.Run("return Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		GetPlayer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()

		GetPlayer(response, request)

		got := response.Body.String()
		want := "10"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}
