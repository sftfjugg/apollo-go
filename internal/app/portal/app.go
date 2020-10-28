package portal

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/services"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/app"
	"go.didapinche.com/foundation/apollo-plus/internal/pkg/http"
	"go.didapinche.com/zeus-go/v2/server"
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

// NewApp is constructor of App,这里引用addressSerive，使address开始去拉取对应adminservice的IP地址
func NewApp(o *Options, logger *zap.Logger, hs *http.Server, zs *server.Server, addr *services.AddressService) (*app.Application, error) {
	a, err := app.New(o.Name, logger,
		app.HTTPServerOption(hs), app.ZeusServerOption(zs))

	if err != nil {
		return nil, errors.Wrap(err, "new app error")
	}

	addr.Poll()
	return a, nil
}

var ProviderSet = wire.NewSet(NewApp, NewOptions)
