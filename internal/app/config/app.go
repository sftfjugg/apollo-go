package config

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.didapinche.com/foundation/apollo-plus/internal/app/config/services"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/app"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/http"
	"go.uber.org/zap"
)

// Options define options
type Options struct {
	Id string
}

// NewOptions is constructor of App options
func NewOptions(v *viper.Viper, logger *zap.Logger, release services.ReleaseMessageService) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal app option error")
	}

	logger.Info("load application options success")

	release.Poll()

	return o, err
}

// NewApp is constructor of App
func NewApp(o *Options, logger *zap.Logger, hs *http.Server) (*app.Application, error) {
	a, err := app.New(o.Id, logger,
		app.HTTPServerOption(hs))

	if err != nil {
		return nil, errors.Wrap(err, "new app error")
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp, NewOptions)
