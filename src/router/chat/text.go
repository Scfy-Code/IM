package chat

import (
	"net/http"

	"github.com/Scfy-Code/scfy-im/app"
)

type TextMessage struct {
}

func (tm TextMessage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := app.NewSession(w, r)
	if !session.IsExist("USER") {
		app.WarnLogger.Println("用户不存在!")
	}

}
func NewSendMessage() *TextMessage {
	return &TextMessage{}
}
