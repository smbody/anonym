package utils

import "github.com/google/uuid"

func AnonymKey() string {
	return uuid.New().String()
}
