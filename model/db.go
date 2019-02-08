package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kelseyhightower/envconfig"
	"github.com/weet/config"
)

var db = NewDBConn()

func NewDBConn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	return db
}

func GetDBConn() *gorm.DB {
	return db
}

func GetDBConfig() (string, string) {
	var config config.Config
	envconfig.Process("", &config)
	mysql := config.MySQL.User + ":" + config.MySQL.Pass + "@tcp(" + config.MySQL.Port + ")/" + config.MySQL.Database + "?charset=utf8&parseTime=True&loc=Local"
	return "mysql", mysql
}
