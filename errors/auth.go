package errors

const (
	errorGeneratingToken = 401
	secretNotValid       = 405
	tokenNotValid        = 404
)

func ErrorGeneratingToken() {
	throw(ServerError{
		ErrCode: errorGeneratingToken,
		Status:  500,
		ErrText: "Error generating token",
	})
}

func SecretNotValid() {
	throw(ServerError{
		ErrCode: secretNotValid,
		Status:  401,
		ErrText: "Secret validation error",
	})
}

func TokenNotValid() {
	throw(ServerError{
		ErrCode: tokenNotValid,
		Status:  401,
		ErrText: "Token validation error",
	})
}
