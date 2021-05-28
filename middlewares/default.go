package middlewares

import (
	"github.com/justinas/alice"
	"net/http"
)

var defaultHandler = alice.New(errorHandling, cors)

func Route(handler http.HandlerFunc) http.Handler {
	return defaultHandler.Then(handler)
}
