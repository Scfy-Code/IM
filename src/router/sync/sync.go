package sync

import (
	"encoding/json"
	"fmt"

	"github.com/Scfy-Code/scfy-im/app"
	"github.com/Scfy-Code/scfy-im/entity"
	"golang.org/x/net/websocket"
)

// MessageChannel 消息通道
var MessageChannel = make(chan *entity.Msg, 1000)

// ConnPool 连接池
var ConnPool = make(map[string]*websocket.Conn)

// Message 消息实体
type Message struct {
	Sender   string
	Receiver string
	Content  string
	MsgType  string
}

// CreateConn 与客户端建立长连接并将连接保存在连接池中
func CreateConn(conn *websocket.Conn) {
	//1、获取请求体
	r := conn.Request()
	//2、解析请求体
	session := app.NewSession(nil, r)
	if !session.IsExist("USER") {
		conn.Close()
		return
	}
	user := session.GetData("USER")
	ConnPool[fmt.Sprintf("%s", user["id"])] = conn
}

// SendMessage 使用websocket发送消息
func SendMessage() {
	go func() {
		select {
		case msg := <-MessageChannel:
			data, err0 := json.Marshal(msg)
			if err0 != nil {
				app.WarnLogger.Printf("解析消息出错！错误信息：%s", err0.Error())
			}
			err := websocket.Message.Send(ConnPool[""], data)
			if err != nil {
				ConnPool["msg.Receiver"].Close()
				delete(ConnPool, "msg.Receiver")
			}
		default:
			return
		}
	}()
}
