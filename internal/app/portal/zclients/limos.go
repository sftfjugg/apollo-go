package zclients

import (
	"github.com/pkg/errors"
	"go.didapinche.com/goapi/apollo_thrift_service/v2"
	"go.didapinche.com/goapi/plat_limos_rpc"
	"go.didapinche.com/zeus-go/v2"
	"go.didapinche.com/zeus-go/v2/client"
)

func NewLimosService(z *zeus.Zeus) (plat_limos_rpc.TChanLimosService, error) {
	c, err := client.New(z, "LimosService")
	if err != nil {
		return nil, errors.Wrap(err, "create zclients UicService error")
	}
	tc := plat_limos_rpc.NewTChanLimosServiceClient(c)
	return tc, nil

}

func NewApolloService(z *zeus.Zeus) (apollo_thrift_service.TChanApolloThriftService, error) {
	c, err := client.New(z, "ApolloThriftService")
	if err != nil {
		return nil, errors.Wrap(err, "create zclients UicService error")
	}
	apollo := apollo_thrift_service.NewTChanApolloThriftServiceClient(c)
	return apollo, nil
}
