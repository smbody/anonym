package errors

import (
	"fmt"
	"log"
)

const (
	unknownError      ErrorCode = 101
	notServed         ErrorCode = 102
	notImplementedYet ErrorCode = 103

	errorGeneratingToken = 401
	tokenRequired        = 402
	forbidden            = 403
	tokenNotValid        = 404
	methodNotAllowed     = 405
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
