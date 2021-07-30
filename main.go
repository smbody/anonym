package main

import (
	"fmt"
	"github.com/kabukky/httpscerts"
	"github.com/smbody/anonym/auth"
	"github.com/smbody/anonym/config"
	"github.com/smbody/anonym/middlewares"
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
	log.Println("Server started (Addr = " +listen + ")")

	// Проверяем, доступен ли cert файл.
	log.Println("Check check self signed certificate")
	err := httpscerts.Check("cert.pem", "key.pem")
	// Если он недоступен, то генерируем новый.
	if err != nil {
		err = httpscerts.Generate("cert.pem", "key.pem", listen)
		log.Println("Generate self signed certificate")
		if err != nil {
			log.Println("Error on generate self signed certificate.")
		}
	}
	log.Println("ListenAndServeTLS " + listen )

	if err := http.ListenAndServeTLS(listen, "cert.pem","key.pem",nil); err != nil {
		log.Println(err)
	}
}
