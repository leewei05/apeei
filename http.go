package apeei

import (
	"fmt"
	"net/http"
)

func StreamServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "sc2")
}
