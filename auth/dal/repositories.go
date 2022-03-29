package dal

import (
	"itsln.com/anonym/auth/dal/gorm"
	"itsln.com/anonym/auth/dal/mongo"
	"itsln.com/anonym/config"
	"itsln.com/anonym/model"
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
