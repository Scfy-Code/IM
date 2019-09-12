package chat

import (
	"fmt"
	"net/http"

	"github.com/Scfy-Code/scfy-im/kit"

	"github.com/Scfy-Code/scfy-im/router/sync"

	"github.com/Scfy-Code/scfy-im/entity"

	"github.com/Scfy-Code/scfy-im/app"
)

// 图片消息结构体
type ImageMessage struct {
}

func (im ImageMessage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := app.NewSession(w, r)
	if !session.IsExist("USER") {
		app.WarnLogger.Println("用户不存在!")
	}
	sender := fmt.Sprintf("%s", session.GetData("USER")["id"])
	err0 := r.ParseForm()
	if err0 != nil {
		app.WarnLogger.Printf("解析请求体失败！错误信息：%s", err0.Error())
	}
	receiver := r.PostFormValue("receiver")
	image, imageHeader, err1 := r.FormFile("image")
	if err1 != nil {
		app.WarnLogger.Printf("获取图片失败！错误消息：%s", err0.Error())
	}
	content := kit.Upload(image, imageHeader)
	msg := entity.NewImageMsg(sender, receiver, content)
	sync.MessageChannel <- msg
}

// NewImageMessage 创建一个图片消息对象
func NewImageMessage() *ImageMessage {
	return &ImageMessage{}
}
