package apeei

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETStreams(t *testing.T) {
	t.Run("get sc2 stream", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/streams/sc2", nil)
		response := httptest.NewRecorder()

		server := NewStreamServer()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "sc2"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("get lol stream", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/streams/lol", nil)
		response := httptest.NewRecorder()

		server := NewStreamServer()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "lol"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
