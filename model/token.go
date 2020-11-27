package model

import (
	"crypto/rand"
	"encoding/hex"
)

type Token struct {
	Key string
}

func NewToken() (token *Token, err error) {
	if key, err := randomKey(32); err ==nil {token = &Token{key}}
	return
}

func randomKey(length int) (string, error) {
	b := make([]byte, length)
	 _, err := rand.Read(b)
	return hex.EncodeToString(b), err
}
