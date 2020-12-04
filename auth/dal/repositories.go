package dal

import (
	dal "github.com/smbody/anonym/auth/dal/mongo"
	"github.com/smbody/anonym/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	monga *mongo.Database
	Users UsersRepo
}

type UsersRepo interface {
	Add() *model.User
	FindById(id string) *model.User
}

func Init() *Repositories {
	db := dal.InitMonga()
	return &Repositories{
		db,
		dal.InitUsers(db),
	}
}
