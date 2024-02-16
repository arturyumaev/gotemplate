package main

import (
	"os"

	"github.com/arturyumaev/gotemplate/internal/application"
	"github.com/arturyumaev/gotemplate/version"
)

func main() {
	name := version.APIName
	version := version.APIVersion

	handler := application.NewHTTPHandler(name, version)

	port := os.Getenv("APPLICATION_PORT")

	app := application.NewApplication()
	app.Name = name
	app.Version = version
	app.Port = port
	app.RegisterHTTPHandler(handler)
	app.Run()
}
