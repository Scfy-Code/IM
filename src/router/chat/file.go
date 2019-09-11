package chat

import (
	"net/http"

	"github.com/Scfy-Code/scfy-im/app"
)

type FileMessage struct {
}

func (fm FileMessage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := app.NewSession(w, r)
	if !session.IsExist("USER") {
		app.WarnLogger.Println("用户不存在!")
	}

}
func NewFileMessage() *FileMessage {
	return &FileMessage{}
}
