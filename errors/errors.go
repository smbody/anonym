package errors

import (
	"fmt"
)

type ErrorCode int

type ServerError struct {
	ErrCode ErrorCode
	Status  int
	ErrText string
}

var predefinedErrors = make(map[ErrorCode]*ServerError)

const (
	UnknownError            ErrorCode = 101
	CantDecodeData                    = 102
	CantEncodeData                    = 103
	ErrorGeneratingToken              = 104
	DataSourceError                   = 105
	CantConnectToToDatabase           = 106
	UnknownDatabase                   = 107
	UnknownUser                       = 1001
	UnknownUserId                     = 1002
	InvalidToken                      = 1001
)

func init() {
	PredefineError(UnknownError, 500, "Internal server error")
	PredefineError(CantDecodeData, 400, "Can't decode data")
	PredefineError(CantEncodeData, 500, "Can't marshal object")
	PredefineError(ErrorGeneratingToken, 500, "Error generating token")
	PredefineError(CantConnectToToDatabase, 500, "Database connection error")
	PredefineError(DataSourceError, 500, "Error reading data from database")
	PredefineError(UnknownDatabase, 500, "Unknown database to connect")
	PredefineError(UnknownUser, 500, "Unknown user")
	PredefineError(UnknownUserId, 500, "Unknown user id")
	PredefineError(InvalidToken, 500, "Invalid token")
}

func PredefineError(errCode ErrorCode, status int, errText string) {
	predefinedErrors[errCode] = &ServerError{errCode, status, errText}
}

func GetError(errCode ErrorCode) *ServerError {
	return predefinedErrors[errCode]
}

func Throw(errCode ErrorCode) {
	panic(GetError(errCode)) //need to get details from source error
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("%v: %s", e.ErrCode, e.ErrText)
}
