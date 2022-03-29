package usecase

import (
	"itsln.com/anonym/auth/dal"
	"itsln.com/anonym/errors"
	"itsln.com/anonym/model"
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
	if len(Secret) > 0 {
		if user := l.repo.FindByKey(Secret); user != nil {
			return l.login(user)
		}
	}
	errors.SecretNotValid()
	return nil
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
