package utils

import "github.com/google/uuid"

func AnonymId () string {
	return uuid.New().String()
}
