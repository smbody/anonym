package model

import (
	"itsln.com/anonym/model"
	"strconv"
)

type User struct {
	ID     uint
	Secret string `gorm:"type:varchar(36);not null;uniqueIndex"`
}

func (u User) ToModel() *model.User {
	return &model.User{
		Id:     strconv.FormatUint(uint64(u.ID), 10),
		Secret: u.Secret,
	}
}
