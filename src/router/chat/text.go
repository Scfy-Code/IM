package chat

import (
	"fmt"
	"net/http"

	"github.com/Scfy-Code/scfy-im/router/sync"

	"github.com/Scfy-Code/scfy-im/entity"

	"github.com/Scfy-Code/scfy-im/app"
)

// TextMessage 文本消息处理器
type textMessage struct {
}

func (tm textMessage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := app.NewSession(w, r)
	if !session.IsExist("USER") {
		app.WarnLogger.Println("用户不存在!")
	}
	sender := fmt.Sprintf("%s", session.GetData("USER")["id"])
	err0 := r.ParseForm()
	if err0 != nil {
		app.WarnLogger.Printf("解析请求体出错！错误信息：%s", err0.Error())
	}
	receiver := r.PostFormValue("receiver")
	content := r.PostFormValue("content")
	msg := entity.NewTextMsg(sender, receiver, content)
	// 1、将消息存储数据库
	//2、将消息存入消息队列
	sync.MessageChannel <- msg
}

// NewTextMessage 创建一个文本消息处理器
func NewTextMessage() http.Handler {
	return &textMessage{}
}
