package server

import "net/http"

// ping returns a "pong" message
func ping(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("pong"))
}
