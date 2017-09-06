package mysql

import (
	"fxlibraries/loggers"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"errors"
	"fmt"
	"time"
)

type DBPoolConfig struct {
	Host         string
	Port         int
	User         string
	DBName       string
	Password     string
	Debug        bool
	MaxIdleConns int
	MaxOpenConns int
}

type DBPool struct {
	DB *gorm.DB
}

func NewPool(conf DBPoolConfig) *DBPool {
	if conf.Host == "" || conf.User == "" || conf.DBName == "" || conf.Password == "" || conf.Port <= 0 {
		panic(errors.New("Mysql config error"))
	}
	if conf.MaxIdleConns == 0 {
		conf.MaxIdleConns = 4
	}
	if conf.MaxOpenConns == 0 {
		conf.MaxOpenConns = 4
	}

	var (
		db  *gorm.DB
		err error
	)

	for retry := 0; retry <= 5; retry++ {
		connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			conf.User,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.DBName)
		db, err = gorm.Open("mysql", connStr)
		if err != nil {
			loggers.Warn.Printf("NewDBPool connect to db %s:%d %s error:%s", conf.Host, conf.Port, conf.DBName, err.Error())
			time.Sleep(2 * time.Second)
			loggers.Warn.Printf("NewDBPool Retrying to connect ...")
			continue
		}
		db.LogMode(conf.Debug)
		db.DB().SetMaxIdleConns(conf.MaxIdleConns)
		db.DB().SetMaxOpenConns(conf.MaxOpenConns)
		return &DBPool{db}
	}
	panic(err)
}

func (self *DBPool) NewConn() *gorm.DB {
	return self.DB.New()
}

func (self *DBPool) Close() {
	self.DB.Close()
}
