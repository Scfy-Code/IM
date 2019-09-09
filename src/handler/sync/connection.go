package sync

import (
	"github.com/Scfy-Code/scfy-im/logger"
	"github.com/Scfy-Code/scfy-im/scope"

	"golang.org/x/net/websocket"
)

// Talk 会话方法
func Talk(conn *websocket.Conn) {
	//1、获取请求体
	r := conn.Request()
	//2、解析请求体
	cookie1, err1 := r.Cookie("SESSIONID")
	cookie2, err2 := r.Cookie("ACCOUNT")
	if err1 != nil || err2 != nil {
		conn.Close()
		return
	}
	sessionID := cookie1.Value
	logger.InfoPrintln(sessionID)
	account := cookie2.Value
	//3、验证账户的信息
	//4、存入域中
	scope.MsgChannel[account]["CONN"] = conn
	scope.MsgChannel[account]["MSG"] = make(chan map[string]interface{})
}
