package apeei

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETStream(t *testing.T) {
	t.Run("get sc2 stream", func(t *testing.T) {
		request := newGetStreamRequest("sc2")
		response := httptest.NewRecorder()

		server := NewStreamServer()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "sc2"

		assertBody(t, got, want)
	})

	t.Run("get lol stream", func(t *testing.T) {
		request := newGetStreamRequest("lol")
		response := httptest.NewRecorder()

		server := NewStreamServer()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "lol"

		assertBody(t, got, want)
	})
}

func newGetStreamRequest(streamName string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/streams/"+streamName, nil)
	return request
}

func assertBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
