package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arturyumaev/gotemplate/internal/server"
)

type Application struct {
	httpServer *http.Server
}

func NewApplication() *Application {
	ctx := context.Background()
	app := &Application{}

	handler := server.NewHandler()
	app.httpServer = &http.Server{
		Addr:    ":3000",
		Handler: handler,
	}
	go app.handleGracefulShutdown(ctx)

	return app
}

func (app *Application) Run() {
	log.Printf("listening on %s\n", app.httpServer.Addr)

	err := app.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}
}

func (app *Application) handleGracefulShutdown(ctx context.Context) {
	// os.Interrupt for CTRL+C, SIGTERM for docker container default signal
	signalCtx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// wait for parent or signal context to cancel
	<-signalCtx.Done()

	// make a new context for the shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.httpServer.Shutdown(shutdownCtx); err != nil {
		fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
	}
}
