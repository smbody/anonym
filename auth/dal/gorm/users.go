package gorm

import (
	g "github.com/smbody/anonym/auth/dal/gorm/model"
	"github.com/smbody/anonym/auth/dal/utils"
	"github.com/smbody/anonym/errors"
	"github.com/smbody/anonym/model"
	"gorm.io/gorm"
)

type Users struct {
	db *gorm.DB
}

func (u Users) Add() *model.User {
	newUser:=&g.User{UserId: utils.AnonymId()}
	r := u.db.Create(newUser)
	if r.Error !=nil {errors.Throw(errors.DataSourceError)}
	return newUser.ToModel()
}

func (u Users) FindById(id string) *model.User {
	anm := &g.User{}
	r := u.db.Where(&g.User{UserId:id}).First(anm)
	if r.Error !=nil {errors.Throw(errors.DataSourceError)}
	return anm.ToModel()

}

func initUsers(db *gorm.DB) *Users {
	if err := db.AutoMigrate(&g.User{}); err!=nil {errors.Throw(errors.MigrationError)}
	return &Users{db: db}

}
