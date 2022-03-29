package gorm

import (
	"gorm.io/gorm"
	g "itsln.com/anonym/auth/dal/gorm/model"
	"itsln.com/anonym/auth/dal/utils"
	"itsln.com/anonym/errors"
	"itsln.com/anonym/model"
)

type Users struct {
	db *gorm.DB
}

func initUsers(db *gorm.DB) *Users {
	if err := db.AutoMigrate(&g.User{}); err != nil {
		errors.DatabaseError(err)
	}
	return &Users{db: db}
}

func (u Users) Add() *model.User {
	newUser := &g.User{Secret: utils.AnonymKey()}
	r := u.db.Create(newUser)
	if r.Error != nil {
		errors.WrongData(r.Error.Error())
	}
	return newUser.ToModel()
}

func (u Users) FindByKey(secret string) *model.User {
	anm := &g.User{}

	// NOTE When querying with struct, GORM will only query
	//with non-zero fields, that means if your field’s value
	//is 0, '', false or other zero values, it won’t be used
	//to build query conditions
	r := u.db.Where(&g.User{Secret: secret}).First(anm)
	if r.Error != nil {
		errors.WrongData(r.Error.Error())
	}
	return anm.ToModel()

}
