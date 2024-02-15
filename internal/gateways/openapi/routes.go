package openapi

import "net/http"

func AddRoutes(
	mux *http.ServeMux,
	serviceName string,
	serviceVersion string,
) {
	mux.HandleFunc("/healthz", healthz())
	mux.HandleFunc("/readyz", readyz())
	mux.HandleFunc("/status", status(serviceName, serviceVersion))
}
