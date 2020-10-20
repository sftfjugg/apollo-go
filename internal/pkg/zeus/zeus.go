package zeus

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.didapinche.com/zeus-go/v2"
)

// New is constructor of clients
func New(v *viper.Viper) (*zeus.Zeus, error) {
	appName := v.GetString("app.name")
	z, err := zeus.New(appName)
	if err != nil {
		return nil, errors.Wrap(err, "create clients error")
	}

	return z, nil
}

// ProviderSet is provider set for wire
var ProviderSet = wire.NewSet(New)
