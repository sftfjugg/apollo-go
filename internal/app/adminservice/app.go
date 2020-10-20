package adminservice

import (
	"apollo-adminserivce/internal/pkg/app"
	"apollo-adminserivce/internal/pkg/http"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Options define options
type Options struct {
	Name string
}

// NewOptions is constructor of App options
func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal app option error")
	}

	logger.Info("load application options success")

	return o, err
}

// NewApp is constructor of App
func NewApp(o *Options, logger *zap.Logger, hs *http.Server) (*app.Application, error) {
	a, err := app.New(o.Name, logger,
		app.HTTPServerOption(hs))

	if err != nil {
		return nil, errors.Wrap(err, "new app error")
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp, NewOptions)