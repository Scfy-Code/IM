package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Scfy-Code/scfy-im/handler"
	"golang.org/x/net/websocket"
)

// Talk 会话方法
func Talk(conn *websocket.Conn) {
	var message string
	for {
		err0 := websocket.Message.Receive(conn, &message)
		if err0 == nil {
			log.Printf("收到消息:%s", message)
		} else {
			log.Printf("消息发送失败！错误信息：%s", err0.Error())
			break
		}
		var sendMessage = map[string]interface{}{"id": "456", "content": "你好世界", "time": time.Now().Unix(), "messageType": "text"}
		data, _ := json.Marshal(sendMessage)
		err1 := websocket.Message.Send(conn, string(data))
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
	http.Handle("/", handler.IndexHandler{"../lib/views/index.scfy"})
	http.Handle("/index.scfy", handler.IndexHandler{"../lib/views/index.scfy"})
}
func init() {
	http.Handle("/talk.action", websocket.Handler(Talk))
}
