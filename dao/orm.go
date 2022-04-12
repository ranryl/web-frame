package dao

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once

func GetDB() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(mysql.Open(viper.GetString("dsn")), &gorm.Config{})
		if err != nil {
			logrus.Panicf("commect db error: ", err)
		}
	})
	return db
}
