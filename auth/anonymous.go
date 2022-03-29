package auth

import (
	"net/http"

	"itsln.com/anonym/auth/dal"
	"itsln.com/anonym/auth/dts"
	"itsln.com/anonym/auth/usecase"
)

type Anonymous interface {
	SignUp(writer http.ResponseWriter, request *http.Request)
	SignIn(writer http.ResponseWriter, request *http.Request)
	Verify(writer http.ResponseWriter, request *http.Request)
}

func Init() Anonymous {
	return dts.Init(usecase.Init(dal.Init()))
}
