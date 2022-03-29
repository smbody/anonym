package auth

import (
	"net/http"

	"github.com/smbody/anonym/auth/dal"
	"github.com/smbody/anonym/auth/dts"
	"github.com/smbody/anonym/auth/usecase"
)

type Anonymous interface {
	SignUp(writer http.ResponseWriter, request *http.Request)
	SignIn(writer http.ResponseWriter, request *http.Request)
	Verify(writer http.ResponseWriter, request *http.Request)
}

func Init() Anonymous {
	return dts.Init(usecase.Init(dal.Init()))
}
