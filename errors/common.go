package errors

import (
	"fmt"
	"log"
)

type ErrorCode int

type ServerError struct {
	ErrCode ErrorCode
	Status  int
	ErrText string
}

func throw(serverErr ServerError) {
	log.Printf("[ERROR] Exception raised: %s", serverErr.Error())
	panic(serverErr)
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("%v - %s", e.ErrCode, e.ErrText)
}
