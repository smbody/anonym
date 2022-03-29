package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"itsln.com/anonym/model"
)

type User struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	Secret string             `bson:"secret,omitempty"`
}

func (u User) ToModel() *model.User {
	return &model.User{
		Id:     u.Id.Hex(),
		Secret: u.Secret,
	}
}
