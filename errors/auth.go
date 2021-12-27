package errors

const (
	errorGeneratingToken = 401
	tokenNotValid        = 404
)

func ErrorGeneratingToken() {
	throw(ServerError{
		ErrCode: errorGeneratingToken,
		Status:  500,
		ErrText: "Error generating token",
	})
}

func TokenNotValid() {
	throw(ServerError{
		ErrCode: tokenNotValid,
		Status:  401,
		ErrText: "Token validation error",
	})
}
