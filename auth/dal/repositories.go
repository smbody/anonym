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
	FindById(id string) *model.User
}

func Init() UsersRepo {
	/*	dsn := config.DataSourceName()
		index := strings.IndexAny(dsn, ":")
		if index>0 && dsn[0:index]=="mongodb" {return mongo.Init()}
	*/
	log.Printf("trying connecting to %s", config.Database())
	if config.Database() == "mongo" {
		return mongo.Init()
	}
	return gorm.Init()
}
