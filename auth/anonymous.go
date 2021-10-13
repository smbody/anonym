package auth

import (
	"github.com/smbody/anonym/auth/dal"
	"github.com/smbody/anonym/auth/dts"
	"github.com/smbody/anonym/auth/usecase"
	"net/http"
)

type Anonymous interface {
	SignUp(writer http.ResponseWriter, request *http.Request)
	SignIn(writer http.ResponseWriter, request *http.Request)
	Verify(writer http.ResponseWriter, request *http.Request)
}

func Init() Anonymous {
	return dts.Init(usecase.Init(dal.Init()))
}
