package config

import (
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// New define constructor *viper.Viper
func New(path string) (*viper.Viper, error) {
	var (
		err error
		v   = viper.New()
	)
	v.AddConfigPath(".")
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err == nil {
		fmt.Printf("use config file -> %s\n", v.ConfigFileUsed())
	} else {
		return nil, err
	}

	return v, err
}

// ProviderSet is wire provider set of config
var ProviderSet = wire.NewSet(New)
