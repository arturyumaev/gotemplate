package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/arturyumaev/gotemplate/internal/server"
)

type Application struct {
	ctx        context.Context
	cancel     context.CancelFunc
	httpServer *http.Server
	exit       *sync.WaitGroup
}

func NewApplication() *Application {
	ctx, cancel := context.WithCancel(context.Background())
	app := &Application{
		ctx:    ctx,
		cancel: cancel,
		exit:   &sync.WaitGroup{},
	}

	handler := server.NewHandler()
	app.httpServer = &http.Server{
		Addr:    ":3000",
		Handler: handler,
	}
	go app.handleGracefulShutdown()

	return app
}

func (app *Application) Run() {
	app.exit.Add(1)
	log.Printf("listening on %s\n", app.httpServer.Addr)

	err := app.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}

	app.cancel()
	app.exit.Wait()
}

func (app *Application) handleGracefulShutdown() {
	defer app.exit.Done()

	// os.Interrupt for CTRL+C, SIGTERM for docker container default signal
	signalCtx, cancel := signal.NotifyContext(app.ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// wait for parent or signal context to cancel
	<-signalCtx.Done()
	fmt.Fprintln(os.Stderr, "gracefully shutting down http server...")

	// make a new context for the shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.httpServer.Shutdown(shutdownCtx); err != nil {
		fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
	}
}
