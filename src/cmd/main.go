package main

import (
	"log"
	"net/http"
	"os"

	indexController "github.com/Scfy-Code/scfy-im/handler/index"
	"github.com/Scfy-Code/scfy-im/handler/sync"
	userController "github.com/Scfy-Code/scfy-im/handler/users"
	"golang.org/x/net/websocket"
)

func main() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
}
func init() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../lib/statics/"))))
	http.Handle("/index.scfy", indexController.IndexView)
	http.Handle("/users/sign_in.scfy", userController.SigninView)
	http.Handle("/users/sign_in.action", userController.SigninAction)
	http.Handle("/users/sign_up.scfy", userController.SignupView)
	http.Handle("/users/sign_up.action", userController.SignupAction)
	http.Handle("/talk.action", websocket.Handler(sync.Talk))
}
