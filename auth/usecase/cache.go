package usecase

import (
	"github.com/smbody/anonym/auth/usecase/redis"
	"github.com/smbody/anonym/model"
)

type Cache interface {
	Find(*model.Token) (*model.Anonym, error)
	Add(token *model.Token, user *model.Anonym) error
}

func InitCache() Cache {
	return redis.InitTokenCache()
}
