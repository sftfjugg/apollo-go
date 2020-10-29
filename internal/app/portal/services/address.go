package services

import (
	"encoding/json"
	"fmt"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/address"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/single_queue"
	"io/ioutil"
	"math/rand"
	http2 "net/http"
	"time"
)

type AddressService struct {
	meta *address.Meta
}

func NewAddress(meta *address.Meta) *AddressService {
	return &AddressService{meta: meta}
}

//一直维护更新ip列表，开启即运行
func (s AddressService) Poll() {
	s.GetAddress("ONLINE", s.meta.ONLINE)
	s.GetAddress("M6V", s.meta.M6V)
	s.GetAddress("ALIYUN", s.meta.ALIYUN)
	s.GetAddress("TEST", s.meta.TEST)
	ticker := time.NewTicker(300 * time.Second)

	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for range ticker.C {
			s.GetAddress("ONLINE", s.meta.ONLINE)
			s.GetAddress("M6V", s.meta.M6V)
			s.GetAddress("ALIYUN", s.meta.ALIYUN)
			s.GetAddress("TEST", s.meta.TEST)
		}
	}(ticker)

}

func (s AddressService) GetAddress(name string, metas []string) {
	i := rand.Intn(len(metas))
	resp, err := http2.Get(metas[i] + "/services/admin")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	address := make([]*models.Address, 0)
	if err := json.Unmarshal(body, &address); err != nil {
		fmt.Println(err, "json.Unmarshal config  failed")
	}
	m := single_queue.GetV()
	if address != nil {
		m[name] = address
	} else {
		fmt.Errorf("get admin by consul error,admin ip no change")
	}
	add, err := json.Marshal(address)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ip list update" + name + string(add))
}
