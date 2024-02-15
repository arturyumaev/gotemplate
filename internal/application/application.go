package application

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arturyumaev/gotemplate/internal/gateways/openapi"
	"github.com/arturyumaev/gotemplate/version"
)

type Application struct {
	ctx    context.Context
	cancel context.CancelFunc
	exit   chan bool

	name       string
	version    string
	httpServer *http.Server
	logger     *slog.Logger
}

func NewApplication(name, version string) *Application {
	ctx, cancel := context.WithCancel(context.Background())

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := &Application{
		ctx:     ctx,
		cancel:  cancel,
		exit:    make(chan bool),
		logger:  logger,
		name:    name,
		version: version,
	}

	mux := http.NewServeMux()
	openapi.AddRoutes(mux, name, version)

	app.httpServer = &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	return app
}

func (app *Application) Run() {
	app.logger.Info(
		"application started",
		"port", app.httpServer.Addr,
		"name", app.name,
		"version", app.version,
		"commit", version.Commit,
		"buildTime", version.BuildTime,
	)

	go app.gracefulShutdown()

	err := app.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		app.logger.Error(
			"error listening and serving",
			"error", err,
		)
	}

	app.cancel()
	<-app.exit
}

func (app *Application) gracefulShutdown() {
	// SIGTERM for docker container default signal
	signalCtx, cancel := signal.NotifyContext(app.ctx, os.Interrupt, os.Kill, syscall.SIGTERM)
	defer cancel()

	// wait for parent or signal context to cancel
	<-signalCtx.Done()
	app.logger.Info("shutting down http server...")

	// make a new context for the shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.httpServer.Shutdown(shutdownCtx); err != nil {
		app.logger.Error(
			"error shutting down http server",
			"error", err,
		)
	}

	app.exit <- true
}
