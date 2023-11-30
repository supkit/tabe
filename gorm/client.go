package gorm

import (
	"github.com/supkit/tabe/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

type Options struct {
	DSN string `yaml:"DSN"`
}

func New(name string, opts ...Option) (*gorm.DB, error) {
	conf, err := config.Watch()
	if err != nil {
		return nil, err
	}

	client := config.GetClientByName(name, conf)
	if strings.Contains(client.Target, "dsn") {
		dsn := strings.Replace(client.Target, "dsn://", "", -1)
		opts = []Option{
			WithDSN(dsn),
		}
	}

	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	return gorm.Open(
		mysql.New(
			mysql.Config{
				DSN: options.DSN,
			},
		),
		&gorm.Config{},
	)
}
