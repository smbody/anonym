package errors

import (
	"fmt"
	"net/url"
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

func MethodNotAllowed(url *url.URL, grantedMethod string, deniedMethod string) {
	throw(ServerError{
		ErrCode: forbidden,
		Status:  500,
		ErrText: fmt.Sprintf("Method %s for %s not allowed. Only %s granted.", deniedMethod, url, grantedMethod),
	})
}
