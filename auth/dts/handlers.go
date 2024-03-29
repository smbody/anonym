package dts

import (
	"itsln.com/anonym/auth/usecase"
	"itsln.com/anonym/model"
	"net/http"
)

type Handlers struct {
	bl *usecase.Logic
}

func Init(l *usecase.Logic) *Handlers {
	return &Handlers{bl: l}
}

func (a *Handlers) SignUp(writer http.ResponseWriter, request *http.Request) {
	writer.Write(Marshal(a.bl.SignUp()))
}

func (a *Handlers) SignIn(writer http.ResponseWriter, request *http.Request) {
	user := model.NewUser()
	Unmarshal(request.Body, user)
	writer.Write(Marshal(a.bl.SignIn(user.Secret)))
}

func (a *Handlers) Verify(writer http.ResponseWriter, request *http.Request) {
	token := &model.Token{}
	Unmarshal(request.Body, token)
	writer.Write(Marshal(a.bl.Verify(token)))
}
