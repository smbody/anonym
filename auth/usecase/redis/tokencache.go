package redis

import (
	"context"
	json "encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/smbody/anonym/config"
	"github.com/smbody/anonym/model"
	"time"
)

type TokenCache struct {
	redis    *redis.Client
	ctx      context.Context
	lifespan time.Duration
}

func InitTokenCache() *TokenCache {
	options := &redis.Options{
		Addr:     config.GetString("redis.server"),
		Password: config.GetString("redis.password"),
		DB:       config.GetInteger("redis.DB"),
	}
	return &TokenCache{
		redis.NewClient(options),
		context.Background(),
		time.Duration(config.GetInteger("token.lifespan")),
	}
}

func (t TokenCache) Find(token *model.Token) (user *model.User, err error) {
	if s, err := t.redis.Get(t.ctx, token.Key).Result(); err ==nil {
		user, err = t.unmarshal([]byte(s))
	}
	return
}

func (t TokenCache) Add(token *model.Token, user *model.User) (err error){
	if s, err := t.marshal(user); err ==nil {
		t.redis.SetEX(t.ctx, token.Key, s, t.lifespan*time.Second)
	}
	return
}

func (t TokenCache) marshal(user *model.User) ([]byte, error) {
	return json.Marshal(user)
}

func (t TokenCache) unmarshal(s []byte) (user *model.User, err error) {
	user = &model.User{}
	err = json.Unmarshal(s, &user)
	return
}
