package zeus

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.didapinche.com/goapi/apollo_thrift_service/v2"
	"go.didapinche.com/zeus-go/v2"
	"go.didapinche.com/zeus-go/v2/server"
)

// New is constructor of zclients
func New(v *viper.Viper) (*zeus.Zeus, error) {
	appName := v.GetString("app.id")
	z, err := zeus.New(appName)
	if err != nil {
		return nil, errors.Wrap(err, "create zclients error")
	}

	return z, nil
}

func NewZeusServer(z *zeus.Zeus, apollo apollo_thrift_service.TChanApolloThriftService) *server.Server {
	ser := server.New(z)
	zService := apollo_thrift_service.NewTChanApolloThriftServiceServer(apollo)
	ser.Register(zService)
	return ser
}

// ProviderSet is provider set for wire
var ProviderSet = wire.NewSet(New, NewZeusServer)
