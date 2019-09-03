package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/Scfy-Code/scfy-im/entry"

	"golang.org/x/net/websocket"
)

// IndexHandler 首页
type IndexHandler struct {
	viewPath string
}

func (ih IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//1、验证是否登录(登录后返回页面/未登录重定向登录页面)
	indexEntry := entry.UserEntry{"123456", "林昊天", "/static/images/avatar.png"}
	//2、
	view, err := template.ParseFiles(ih.viewPath)
	if err == nil {
		view.Execute(w, indexEntry)
	} else {
		log.Println(err.Error())
	}
}

// Talk 会话方法
func Talk(conn *websocket.Conn) {
	var message string
	for {
		err := websocket.Message.Receive(conn, &message)
		if err == nil {
			log.Println(message)
			websocket.Message.Send(conn, "你好世界")
		} else {
			log.Println(err.Error())
		}
	}
}

func main() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
}
func init() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../lib/statics/"))))
	http.Handle("/", IndexHandler{"../lib/views/index.scfy"})
	http.Handle("/index.scfy", IndexHandler{"../lib/views/index.scfy"})
}
func init() {
	http.Handle("/talk.action", websocket.Handler(Talk))
}
