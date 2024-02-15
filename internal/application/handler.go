package application

import (
	"net/http"

	"github.com/arturyumaev/gotemplate/internal/gateways/openapi"
)

func NewHTTPHandler(
	serviceName,
	serviceVersion string,
) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", openapi.Healthz())
	mux.HandleFunc("/readyz", openapi.Readyz())
	mux.HandleFunc("/status", openapi.Status(serviceName, serviceVersion))

	return mux
}
