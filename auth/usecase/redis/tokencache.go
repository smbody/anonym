package redis

import (
	"context"
	json "encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/smbody/anonym/config"
	"github.com/smbody/anonym/model"
	"log"
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
	log.Printf("Connect to Redis. Addr = %s, DB=%v", options.Addr, options.DB)
	return &TokenCache{
		redis.NewClient(options),
		context.Background(),
		time.Duration(config.GetInteger("token.lifespan")),
	}
}

func (t TokenCache) Find(token *model.Token) (user *model.User, err error) {
	s, err := t.redis.Get(t.ctx, token.Key).Result();
	if err == nil {
		user, err = t.unmarshal([]byte(s))
	}
	return
}

func (t TokenCache) Add(token *model.Token, user *model.User) (err error) {
	s, err := t.marshal(user)
	if err == nil {
		err = t.redis.Set(t.ctx, token.Key, s, t.lifespan*time.Second).Err()
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
