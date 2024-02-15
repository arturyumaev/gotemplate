package main

import (
	"net/http"

	"github.com/arturyumaev/gotemplate/internal/application"
	"github.com/arturyumaev/gotemplate/internal/gateways/openapi"
	"github.com/arturyumaev/gotemplate/version"
)

func main() {
	name := version.APIName
	version := version.APIVersion

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", openapi.Healthz())
	mux.HandleFunc("/readyz", openapi.Readyz())
	mux.HandleFunc("/status", openapi.Status(name, version))

	app := application.NewApplication()
	app.Name = name
	app.Version = version
	app.RegisterHTTPHandler(mux)

	app.Run()
}
