package main

import (
	"github.com/arturyumaev/gotemplate/internal/application"
	"github.com/arturyumaev/gotemplate/version"
)

func main() {
	name := version.APIName
	version := version.APIVersion

	handler := application.NewHTTPHandler(name, version)

	app := application.NewApplication()
	app.Name = name
	app.Version = version
	app.RegisterHTTPHandler(handler)

	app.Run()
}
