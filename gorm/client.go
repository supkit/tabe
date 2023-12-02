package gorm

import (
	"github.com/supkit/tabe/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

// DefaultStringSize string 类型字段的默认长度
const DefaultStringSize uint = 256

// DefaultMaxIdleConn 空闲连接数
const DefaultMaxIdleConn int = 10

// DefaultMaxOpenConn 并发连接数
const DefaultMaxOpenConn int = 20

// DefaultConnMaxLifetime 超时时间
const DefaultConnMaxLifetime time.Duration = time.Minute * 10

// New 实例化一个连接
func New(name string, opts ...Option) (*gorm.DB, error) {
	client := config.GetClientByName(name)
	if strings.Contains(client.Target, "dsn") {
		dsn := strings.Replace(client.Target, "dsn://", "", -1)
		opts = []Option{
			WithDSN(dsn),
		}
	}

	options := Options{
		DefaultStringSize: DefaultStringSize,
		MaxIdleConn:       DefaultMaxIdleConn,
		MaxOpenConn:       DefaultMaxOpenConn,
		ConnMaxLifetime:   DefaultConnMaxLifetime,
	}

	for _, o := range opts {
		o(&options)
	}

	mysqlConfig := mysql.Config{
		DSN:                       options.DSN,
		DefaultStringSize:         options.DefaultStringSize,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}

	db, err := gorm.Open(
		mysql.New(mysqlConfig), &gorm.Config{},
	)

	if err != nil {
		return db, err
	}

	conn, err := db.DB()
	if err != nil {
		return db, err
	}

	conn.SetMaxIdleConns(options.MaxIdleConn)
	conn.SetMaxOpenConns(options.MaxIdleConn)
	conn.SetConnMaxLifetime(options.ConnMaxLifetime)

	return db, err
}
