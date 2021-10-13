package gorm

import (
	"github.com/smbody/anonym/model"
	"gorm.io/gorm"
)

type Users struct {
	db *gorm.DB
}

func (u Users) Add() *model.User {
	panic("implement me")
}

func (u Users) FindById(id string) *model.User {
	panic("implement me")
}

func initUsers(db *gorm.DB) *Users {
	return &Users{db: db}

}
