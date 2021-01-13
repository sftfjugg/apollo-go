package single_queue

import (
	"sync"
)

//var address map[string][]*models.Address = make(map[string][]*models.Address)
//
//func GetV() map[string][]*models.Address {
//	return address
//}
var address sync.Map

func GetV() *sync.Map {
	return &address
}
