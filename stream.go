package apeei

type Stream struct{}

func (s *Stream) LiveStreams() []string {

	return []string{"sc2", "lol", "test"}
}
