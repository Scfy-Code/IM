package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Scfy-Code/scfy-im/handler"
	chatController "github.com/Scfy-Code/scfy-im/handler/chat"
	indexController "github.com/Scfy-Code/scfy-im/handler/index"
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
	http.Handle("/index.scfy", indexController.NewIndexView())
	http.Handle("/users/sign_in.scfy", userController.NewSignInView())
	http.Handle("/users/sign_in.action", userController.NewSignInAction())
	http.Handle("/users/sign_up.scfy", userController.NewSignUpView())
	http.Handle("/users/sign_up.action", userController.NewSignInAction())
	http.handle("/chat/send.action", chatController.NewSendMessage())
	http.Handle("/talk.action", websocket.Handler(handler.Talk))
}
