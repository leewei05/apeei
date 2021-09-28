package apeei

import "testing"

func TestGETlivestream(t *testing.T) {
	s := &Stream{}
	got := s.LiveStreams()
	want := []string{"sc2", "lol", "test"}

	if len(got) != len(want) {
		t.Errorf("got %d, want %d", len(got), len(want))
	}
}
