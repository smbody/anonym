package usecase

import (
	"itsln.com/anonym/auth/usecase/redis"
	"itsln.com/anonym/model"
)

type Cache interface {
	Find(*model.Token) (*model.Anonym, error)
	Add(token *model.Token, user *model.Anonym) error
}

func InitCache() Cache {
	return redis.InitTokenCache()
}
