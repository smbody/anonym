package usecase

import (
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


func (l Logic) SignUp() *model.Token {
	user := l.repo.Add()
	return l.login(user)
}

func (l Logic) SignIn(Id string) *model.Token {
	user := l.repo.FindById(Id)
	if user == nil {
		errors.Throw(errors.UnknownUser)
		return nil
	}
	return l.login(user)

}

func (l Logic) login(user *model.User) *model.Token {
	token, err := model.NewToken()
	if err != nil {
		errors.Throw(errors.ErrorGeneratingToken)
	} else {
		l.cache.Add(token, user)
	}
	return token
}

func (l Logic) Verify(token *model.Token) *model.User {
	if user, err := l.cache.Find(token); err !=nil {return user}
	return nil
}


