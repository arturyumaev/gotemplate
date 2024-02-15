package main

import (
	"github.com/arturyumaev/gotemplate/internal/application"
	"github.com/arturyumaev/gotemplate/version"
)

func main() {
	app := application.NewApplication(
		"api",
		version.APIVersion,
	)
	app.Run()
}
