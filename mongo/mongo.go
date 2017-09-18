package mongo

import (
	"errors"
	"fmt"
	"fxlibraries/loggers"
	"log"
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

type MongodbConfig struct {
	Host   string
	Port   int
	DBName string
	Debug  bool
}

type MgoPool struct {
	*mgo.Database
}

const RetryCount = 5

func NewPool(conf *MongodbConfig) *MgoPool {
	if conf.Host == "" || conf.Port <= 0 || conf.DBName == "" {
		panic(errors.New("MongoDB config error"))
	}
	mgoUrl := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	var (
		session *mgo.Session
		err     error
	)
	if conf.Debug {
		mgo.SetDebug(true)
		mLogger := log.New(os.Stderr, "", log.LstdFlags)
		mgo.SetLogger(mLogger)
	}
	for i := 0; i < RetryCount; i++ {
		mgo.SetDebug(conf.Debug)
		session, err = mgo.Dial(mgoUrl)
		if err != nil {
			loggers.Error.Printf("Failed to connect mongodb: %v", conf)
			time.Sleep(2 * time.Second)
			loggers.Warn.Printf("Retrying to connect to mongodb: %v", conf)
			continue
		}
		db := session.DB(conf.DBName)
		return &MgoPool{db}
	}
	panic(err)

}

func NotFound(err error) bool {
	return err == mgo.ErrNotFound
}
