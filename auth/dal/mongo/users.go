package mongo

import (
	"context"
	mongoModel "github.com/smbody/anonym/auth/dal/mongo/model"
	"github.com/smbody/anonym/errors"
	"github.com/smbody/anonym/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users struct {
	ctx      context.Context
	userList *mongo.Collection
}

func initUsers(monga *mongo.Database) *Users {
	return &Users{
		ctx:      context.Background(),
		userList: monga.Collection("Users"),
	}
}

func (r Users) Add() (user *model.User) {
	id := primitive.NewObjectID()
	mongoUser := &mongoModel.User{Id: id, Secret: user.Secret}
	if _, err := r.userList.InsertOne(r.ctx, mongoUser); err != nil {
		errors.DatabaseError(err)
	} else {
		user = mongoUser.ToModel()
	}
	return
}

func (r Users) FindByKey(secret string) (user *model.User) {
	mongoUser := &mongoModel.User{}
	if err := r.userList.FindOne(r.ctx, bson.M{"secret": secret}).Decode(mongoUser); err != nil {
		errors.CantDecodeData(err)
		return nil
	}
	return mongoUser.ToModel()
}
