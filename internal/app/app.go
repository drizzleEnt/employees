package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/drizzleent/emplyees/internal/config"
	"github.com/drizzleent/emplyees/pkg/closer"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDebs(ctx)

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	err := a.runHTTPServer()

	if err != nil {
		log.Fatalf("failed to run http server %s", err.Error())
	}

	return nil
}

func (a *App) initDebs(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = NewServiceProvider()
	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	a.httpServer = &http.Server{
		Addr:           a.serviceProvider.HTTPConfig().Address(),
		Handler:        a.serviceProvider.Handler(ctx).InitRoutes(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return nil
}

func (a *App) runHTTPServer() error {
	log.Printf("HTTP server running on: %s", a.serviceProvider.HTTPConfig().Address())

	err := a.httpServer.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}
