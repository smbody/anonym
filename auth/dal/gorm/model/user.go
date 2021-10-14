package model

import "github.com/smbody/anonym/model"

type User struct {
	ID uint `gorm:"primaryKey"`
	UserId string `gorm:"column:USERID"`
}

func (u User) ToModel() *model.User {
	return &model.User{
		Id: u.UserId,
	}
}
