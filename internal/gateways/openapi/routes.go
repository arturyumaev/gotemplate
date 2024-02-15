package openapi

import "net/http"

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/healthz", healthz())
	mux.HandleFunc("/readyz", readyz())
}
