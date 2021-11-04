package errors

import "fmt"

const (
	unknownDatabase = 200
	databaseError   = 201
	cantDecodeData  = 202
	cantEncodeData  = 203
	wrongId         = 204
	wrongData       = 205
)

func UnknownDatabase(db string) {
	throw(ServerError{
		ErrCode: unknownDatabase,
		Status:  500,
		ErrText: fmt.Sprintf("Unknown database type: %s", db),
	})
}

func DatabaseError(err error) {
	throw(ServerError{
		ErrCode: databaseError,
		Status:  500,
		ErrText: fmt.Sprintf("Error occurred while communication with Database: %s", err.Error()),
	})
}

func CantDecodeData(err error) {
	throw(ServerError{
		ErrCode: cantDecodeData,
		Status:  400,
		ErrText: fmt.Sprintf("Can't decode data: %s", err.Error()),
	})
}

func CantEncodeData(err error) {
	throw(ServerError{
		ErrCode: cantEncodeData,
		Status:  500,
		ErrText: fmt.Sprintf("Can't marshal object: %s", err.Error()),
	})
}

func WrongData(details string) {
	throw(ServerError{
		ErrCode: wrongId,
		Status:  500,
		ErrText: fmt.Sprintf("Wrong data: %s", details),
	})
}

func WrongId() {
	throw(ServerError{
		ErrCode: wrongId,
		Status:  500,
		ErrText: "Wrong id",
	})
}
