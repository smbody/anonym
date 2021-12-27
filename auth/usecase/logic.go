package usecase

import (
	"fmt"
	"github.com/smbody/anonym/auth/dal"
	"github.com/smbody/anonym/errors"
	"github.com/smbody/anonym/model"
)

type Logic struct {
	cache Cache
	repo  dal.UsersRepo
}

func Init(r dal.UsersRepo) *Logic {
	return &Logic{
		cache: InitCache(),
		repo:  r,
	}
}

func (l Logic) SignUp() *model.User {
	user := l.repo.Add()
	return user
}

func (l Logic) SignIn(Secret string) *model.Token {
	user := l.repo.FindByKey(Secret)
	if user == nil {
		errors.WrongData(fmt.Sprintf("Cant find user by id =%s", Secret))
		return nil
	}
	return l.login(user)

}

func (l Logic) login(user *model.User) (token *model.Token) {
	if token, err := model.NewToken(); err == nil {
		if err = l.cache.Add(token, user.ToAnonym()); err == nil {
			return token
		}
	}
	errors.ErrorGeneratingToken()
	return nil
}

func (l Logic) Verify(token *model.Token) *model.Anonym {
	if anm, err := l.cache.Find(token); err == nil {
		return anm
	}
	errors.TokenNotValid()
	return nil
}
