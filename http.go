package apeei

import (
	"fmt"
	"net/http"
	"strings"
)

type StreamServer struct {
	http.Handler
}

func NewStreamServer() *StreamServer {
	s := new(StreamServer)

	router := http.NewServeMux()
	router.HandleFunc("/streams/", StreamHandler)

	s.Handler = router
	return s
}

func StreamHandler(w http.ResponseWriter, r *http.Request) {
	streamName := strings.TrimPrefix(r.URL.Path, "/streams/")

	response := getStreamResponse(streamName)
	fmt.Fprintf(w, response)
}

func getStreamResponse(streamName string) string {
	switch streamName {
	case "sc2":
		return "sc2"
	case "lol":
		return "lol"
	}

	return ""
}
