package services

import (
	"encoding/json"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/address"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/models"
	"go.didapinche.com/foundation/apollo-plus/internal/app/portal/single_queue"
	"go.uber.org/zap"
	"io/ioutil"
	"math/rand"
	http2 "net/http"
	"time"
)

type AddressService struct {
	meta *address.Meta
	log  *zap.Logger
}

func NewAddress(meta *address.Meta, log *zap.Logger) *AddressService {
	return &AddressService{meta: meta, log: log}
}

//一直维护更新ip列表，开启即运行
func (s AddressService) Poll() {
	ticker := time.NewTicker(100 * time.Second)
	s.GetAddress("ONLINE", s.meta.ONLINE)
	s.GetAddress("ALIYUN", s.meta.ALIYUN)
	s.GetAddress("TEST", s.meta.TEST)

	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for range ticker.C {
			s.GetAddress("ONLINE", s.meta.ONLINE)
			s.GetAddress("ALIYUN", s.meta.ALIYUN)
			s.GetAddress("TEST", s.meta.TEST)
		}
	}(ticker)

}

func (s AddressService) GetAddress(name string, metas []string) {
	i := rand.Intn(len(metas))
	resp, err := http2.Get(metas[i] + "/services/admin")
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	address := make([]*models.Address, 0)
	if err := json.Unmarshal(body, &address); err != nil {
		s.log.Error("json.Unmarshal config  failed")
	}
	m := single_queue.GetV()
	if address != nil {
		m[name] = address
	} else {
		s.log.Error("get admin by consul error,admin ip no change")
	}
	add, err := json.Marshal(address)
	if err != nil {
		s.log.Error("json.Unmarshal config  failed")
	}
	s.log.Info("ip list update" + name + string(add))
}
