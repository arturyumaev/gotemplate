package server

import (
	"net/http"

	"github.com/arturyumaev/gotemplate/internal/gateways/openapi"
)

// NewHandler constructor is responsible for all the top-level HTTP stuff that applies to all endpoints
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	openapi.AddRoutes(mux)

	return mux
}
