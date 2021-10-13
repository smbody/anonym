package model

import (
	"github.com/smbody/anonym/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
}

func (u User) ToModel() *model.User {
	return &model.User{
		Id: u.Id.Hex(),
	}
}
