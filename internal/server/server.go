package server

import "net/http"

// NewHandler constructor is responsible for all the top-level HTTP stuff that applies to all endpoints
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	var handler http.Handler = mux
	return handler
}
