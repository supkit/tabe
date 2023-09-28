package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Options struct {
	DSN string `yaml:"DSN"`
}

func New(name string, opts ...Options) (*gorm.DB, error) {
	return gorm.Open(
		mysql.New(
			mysql.Config{
				DSN: "",
			},
		),
	)
}
