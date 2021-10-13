package dal

import (
	"github.com/smbody/anonym/auth/dal/mongo"
	"github.com/smbody/anonym/model"
)


type UsersRepo interface {
	Add() *model.User
	FindById(id string) *model.User
}

func Init() UsersRepo {
	return mongo.Init()
}
