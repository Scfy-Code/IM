package chat

import (
	"log"
	"net/http"
	"time"

	"github.com/Scfy-Code/scfy-im/entry"
)

type SendMessage struct {
}

func (sm SendMessage) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	cookie0, err0 := r.Cookie("SESSIONID")
	cookie1, err1 := r.Cookie("SENDERID")
	if err0 != nil || err1 != nil {
		return
	}
	sessionID := cookie0.Value
	log.Println(sessionID)
	senderID := cookie1.Value
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
	}
	receiverID := r.FormValue("receiverID")
	content := r.FormValue("content")
	//将消息存入数据库中然后将消息存入消息通道
	entry.MessageChannel[receiverID] <- entry.Message{SenderID: senderID, ReceiverID: receiverID, Time: time.Now(), Content: content}
	w.Write([]byte("消息发送成功"))
}
func NewSendMessage() *SendMessage {
	return &SendMessage{}
}
