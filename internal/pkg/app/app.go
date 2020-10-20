package app

import (
	"apollo-adminserivce/internal/pkg/http"
	"github.com/pkg/errors"
	"go.didapinche.com/juno-go/v2"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

// Application define application struct
type Application struct {
	name       string
	logger     *zap.Logger
	httpServer *http.Server
}

// Option define Optional function
type Option func(app *Application) error

// HTTPServerOption returns optional function of address server
func HTTPServerOption(svr *http.Server) Option {
	return func(app *Application) error {
		app.httpServer = svr

		return nil
	}
}

// New is constructor of application
func New(name string, logger *zap.Logger, options ...Option) (*Application, error) {
	app := &Application{
		name:   name,
		logger: logger.With(zap.String("type", "Application")),
	}

	for _, option := range options {
		if err := option(app); err != nil {
			return nil, err
		}
	}

	return app, nil
}

// Start application
func (a *Application) Start() error {
	if a.httpServer != nil {
		if err := a.httpServer.Start(); err != nil {
			return errors.Wrap(err, "address server start error")
		}
	}

	err := juno.RegisterWithParams(&juno.Params{
		Name: a.name,
		Port: a.httpServer.Port(),
		Addr: a.httpServer.Host(),
	})
	if err != nil {
		return errors.Wrap(err, "juno register error")
	}

	return nil
}

// AwaitSignal graceful shutdown when receive terminal signal
// kill -15
func (a *Application) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-c:
		a.logger.Info("receive a signal", zap.String("signal", s.String()))
		if a.httpServer != nil {
			if err := a.httpServer.Stop(); err != nil {
				a.logger.Warn("stop address server error", zap.Error(err))
			}
		}

		os.Exit(0)
	}
}
