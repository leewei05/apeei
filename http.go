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
	streamID := strings.TrimPrefix(r.URL.Path, "/streams/")

	switch streamID {
	case "sc2":
		fmt.Fprintf(w, "sc2")
	case "lol":
		fmt.Fprintf(w, "lol")
	}
}
