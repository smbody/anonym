package main

import (
	"fmt"
	"github.com/smbody/anonym/auth"
	"github.com/smbody/anonym/config"
	"log"
	"net/http"
)

var server = auth.Init()

func main() {
	config.Init()

	//todo обработка ошибок
	http.HandleFunc("/signup", server.SignUp)
	http.HandleFunc("/signin", server.SignIn)
	http.HandleFunc("/verify", server.Verify)

	listen := fmt.Sprintf("%s:%s", config.GetString("host"), config.GetString("port"))
	log.Println("Server started (addr:" +listen + ")")
	if err := http.ListenAndServe(listen, nil); err != nil {
		log.Println(err)
	}
}
