package apeei

import (
	"fmt"
	"net/http"
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
	fmt.Fprintf(w, "sc2")
}
