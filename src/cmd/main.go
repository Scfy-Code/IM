package main

import (
	"net/http"
	"os"

	"github.com/Scfy-Code/scfy-im/handler/index"
	"github.com/Scfy-Code/scfy-im/handler/sync"
	"github.com/Scfy-Code/scfy-im/handler/users"
	"github.com/Scfy-Code/scfy-im/logger"
	"golang.org/x/net/websocket"
)

func main() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.ErrPrintf("端口监听错误！错误原因：%s", err.Error())
		os.Exit(2)
	}
}
func init() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../lib/statics/"))))
	http.Handle("/index.scfy", index.NewIndexController())
	http.Handle("/user/sign_in.scfy", users.NewSigninView())
	http.Handle("/user/sign_in.action", users.NewSigninAction())
	http.Handle("/user/sign_up.scfy", users.NewSignupView())
	http.Handle("/user/sign_up.action", users.NewSignupAction())
	http.Handle("/talk.action", websocket.Handler(sync.Talk))
}
