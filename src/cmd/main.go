package main

import (
	"net/http"
	"os"

	"github.com/Scfy-Code/scfy-im/app"
	"github.com/Scfy-Code/scfy-im/router/chat"
	"github.com/Scfy-Code/scfy-im/router/index"
	"github.com/Scfy-Code/scfy-im/router/sync"
	"github.com/Scfy-Code/scfy-im/router/users"
	"golang.org/x/net/websocket"
)

func main() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		app.ErrorLogger.Printf("端口监听错误！错误原因：%s", err.Error())
		os.Exit(2)
	}
}
func init() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(app.StaticDir))))
	http.Handle("/index.scfy", index.NewIndexController())
	http.Handle("/user/sign_in.scfy", users.NewSigninView())
	http.Handle("/user/sign_in.action", users.NewSigninAction())
	http.Handle("/user/sign_up.scfy", users.NewSignupView())
	http.Handle("/user/sign_up.action", users.NewSignupAction())
	http.Handle("/msg/textMessage.action", chat.NewTextMessage())
	http.Handle("/msg/imageMessage.action", chat.NewImageMessage())
	http.Handle("/msg/fileMessage.action", chat.NewFileMessage())
	http.Handle("/talk.action", websocket.Handler(sync.CreateConn))
}
