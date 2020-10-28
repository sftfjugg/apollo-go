package single_queue

import "go.didapinche.com/foundation/apollo-plus/internal/app/portal/models"

var address map[string][]*models.Address = make(map[string][]*models.Address)

func GetV() map[string][]*models.Address {
	return address
}
