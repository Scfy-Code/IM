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
	indexEntry := &entry.UserEntry{ID: "123456", RemarkName: "林昊天", Avatar: "/static/images/avatar.png"}
	chatList := []map[string]interface{}{{"ID": "77425000", "RemarkName": "汤姆", "Avatar": "/static/images/avatar.png"}}
	friendList := []map[string]interface{}{{"ID": "77425000", "RemarkName": "汤姆", "Avatar": "/static/images/avatar.png"}}
	startList := []map[string]interface{}{{"ID": "77425000", "RemarkName": "汤姆", "Avatar": "/static/images/avatar.png"}}
	indexEntry.ChatList = chatList
	indexEntry.FriendList = friendList
	indexEntry.StartList = startList
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
		err0 := websocket.Message.Receive(conn, &message)
		if err0 == nil {
			log.Printf("收到消息！消息%s", message)
		} else {
			log.Printf("消息发送失败！错误信息：%s", err0.Error())
			break
		}
		err1 := websocket.Message.Send(conn, "你好世界")
		if err1 == nil {
			log.Println("消息发送成功")
		} else {
			log.Printf("消息发送失败！错误信息：%s", err1.Error())
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
