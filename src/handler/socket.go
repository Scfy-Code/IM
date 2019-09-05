package handler

import (
	"encoding/json"
	"log"

	"github.com/Scfy-Code/scfy-im/entry"

	"golang.org/x/net/websocket"
)

// Talk 会话方法
func Talk(conn *websocket.Conn) {
	//1、获取请求体
	r := conn.Request()
	//2、解析请求体
	cookie1, err1 := r.Cookie("SESSIONID")
	cookie2, err2 := r.Cookie("SENDERID")
	if err1 != nil || err2 != nil {
		conn.Close()
		return
	}
	sessionID := cookie1.Value
	log.Println(sessionID)
	senderID := cookie2.Value
	var message string
	for {
		err0 := websocket.Message.Receive(conn, &message)
		if err0 == nil {
			log.Printf("收到消息:%s", message)
		} else {
			log.Printf("消息发送失败！错误信息：%s", err0.Error())
			break
		}
		//监听消息通道
		select {
		case ms := <-entry.MessageChannel[senderID]:
			data, err3 := json.Marshal(ms)
			if err3 != nil {
				continue
			}
			websocket.Message.Send(conn, data)
		}
	}
}
