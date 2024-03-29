package main

import (
	"fmt"
	"itsln.com/anonym/auth"
	"itsln.com/anonym/config"
	"itsln.com/anonym/middlewares"
	"log"
	"net/http"
)

func main() {
	config.Init()
	server := auth.Init()
	http.Handle("/signup", middlewares.Route(server.SignUp))
	http.Handle("/signin", middlewares.Route(server.SignIn))
	http.Handle("/verify", middlewares.Route(server.Verify))
	listen := fmt.Sprintf("%s:%s", config.GetString("host"), config.GetString("port"))
	log.Println("Server started (Addr = " + listen + ")")
	if err := http.ListenAndServe(listen, nil); err != nil {
		log.Println(err)
	}
}
