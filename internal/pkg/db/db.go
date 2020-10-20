package db

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
	"go.uber.org/zap"
)

type Options struct {
	URL   string `yaml:"url"`
	Debug bool
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("db", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal db option error")
	}

	return o, err
}

func New(o *Options, logger *zap.Logger) (*gorm.DB, error) {
	var err error
	db, err := gorm.Open("mysql", o.URL)
	if err != nil {
		return nil, errors.Wrap(err, "gorm open db connection error")
	}

	if o.Debug {
		db = db.Debug()
	}
	return db, nil
}

// ProviderSet define provider set of db
var ProviderSet = wire.NewSet(New, NewOptions)
