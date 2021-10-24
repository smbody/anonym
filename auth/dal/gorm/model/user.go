package model

import (
	"github.com/smbody/anonym/model"
)

type User struct {
	ID     uint
	UserId string `gorm:"type:varchar(36);not null;uniqueIndex"`
}

func (u User) ToModel() *model.User {
	return &model.User{
		Id: u.UserId,
	}
}
