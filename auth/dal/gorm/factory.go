package gorm

import (
	"github.com/smbody/anonym/config"
	"github.com/smbody/anonym/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Init() *Users {
	return initUsers(initGorm())
}

func initGorm() *gorm.DB {
	dsn := config.DataSourceName()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=nil {
		log.Fatalf("Error occurred while establishing connection to %s", dsn)
		errors.Throw(errors.CantConnectToToDatabase)
	}
	return db

}