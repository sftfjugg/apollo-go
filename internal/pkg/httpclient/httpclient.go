package httpclient

import (
	"github.com/google/wire"
	"net/http"
	"time"
)

func New() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	return &http.Client{Transport: tr}
}

var ProviderSet = wire.NewSet(New)
