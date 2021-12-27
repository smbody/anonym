package dal

import (
	"github.com/smbody/anonym/auth/dal/gorm"
	"github.com/smbody/anonym/auth/dal/mongo"
	"github.com/smbody/anonym/config"
	"github.com/smbody/anonym/model"
	"log"
)

type UsersRepo interface {
	Add() *model.User
	FindByKey(Secret string) *model.User
}

func Init() UsersRepo {
	log.Printf("Using %s database", config.Database())
	//log.Printf("uri=%s", config.DataSourceName())
	if config.Database() == "mongo" {
		return mongo.Init()
	}
	return gorm.Init()
}
