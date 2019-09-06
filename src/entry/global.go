package entry

import "golang.org/x/net/websocket"

var (
	// MessageChannel 声明一个消息的通道，所有的消息都通过通道获取
	MessageChannel = make(map[string]chan Message)
	// ConnectionMap 连接存储器存储所有的连接
	ConnectionMap = make(map[string]*websocket.Conn)
	// SessionMap session存储器
	SessionMap = make(map[string]UserEntry)
)
