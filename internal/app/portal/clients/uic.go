package clients

import (
	"github.com/pkg/errors"
	"go.didapinche.com/goapi/uic_service_api"
	"go.didapinche.com/zeus-go/v2"
	"go.didapinche.com/zeus-go/v2/client"
)

// NewUicService is constructor of TChanUicService
func NewUicService(z *zeus.Zeus) (uic_service_api.TChanUicService, error) {
	c, err := client.New(z, "UicService")
	if err != nil {
		return nil, errors.Wrap(err, "create clients UicService error")
	}

	tc := uic_service_api.NewTChanUicServiceClient(c)

	return tc, nil

}
