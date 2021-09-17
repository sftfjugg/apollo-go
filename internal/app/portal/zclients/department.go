package zclients

import (
	"go.didapinche.com/goapi/user_department_service_api"
	"go.didapinche.com/zeus-go/v2"
	"go.didapinche.com/zeus-go/v2/client"
)

func NewDepartmentService(z *zeus.Zeus) (user_department_service_api.TChanDepartmentService, error) {
	c, err := client.New(z, "DepartmentService")
	if err != nil {
		//return nil, errors.Wrap(err, "create zeus LimosService error")
		return nil, nil
	}
	tc := user_department_service_api.NewTChanDepartmentServiceClient(c)
	return tc, nil
}
