package address

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

type Options struct {
	TEST   string `yaml:"TEST"`
	ALIYUN string `yaml:"ALIYUN"`
	ONLINE string `yaml:"ONLINE"`
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("meta", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal db option error")
	}

	return o, err
}

//meta地址
type Meta struct {
	TEST   []string `yaml:"TEST"`
	ALIYUN []string `yaml:"ALIYUN"`
	ONLINE []string `yaml:"ONLINE"`
}

func NewMetas(o *Options) (*Meta, error) {
	var err error
	m := new(Meta)
	m.TEST = strings.Split(o.TEST, ",")
	m.ALIYUN = strings.Split(o.ALIYUN, ",")
	m.ONLINE = strings.Split(o.ONLINE, ",")
	return m, err
}

var ProviderSet = wire.NewSet(NewMetas, NewOptions)
