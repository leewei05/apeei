package main

import (
	"log"
	"net/http"

	"github.com/leewei05/apeei"
)

func main() {
	handler := http.HandlerFunc(apeei.StreamServer)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
