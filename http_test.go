package apeei

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETStreams(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/api/v1/streams", nil)
	response := httptest.NewRecorder()

	StreamServer(response, request)

	got := response.Body.String()
	want := "sc2"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
