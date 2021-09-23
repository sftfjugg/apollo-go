package zclients

import (
	"go.didapinche.com/goapi/user_department_service_api"
	"go.didapinche.com/zeus-go/v2"
	"go.didapinche.com/zeus-go/v2/client"
)

func NewUserService(z *zeus.Zeus) (user_department_service_api.TChanUserService, error) {
	c, err := client.New(z, "UserService")
	if err != nil {
		//return nil, errors.Wrap(err, "create zeus LimosService error")
		return nil, nil
	}
	tc := user_department_service_api.NewTChanUserServiceClient(c)
	return tc, nil
}
