package single_queue

import "apollo-adminserivce/internal/app/portal/models"

var address map[string][]*models.Address = make(map[string][]*models.Address)

func GetV() map[string][]*models.Address {
	return address
}
