package entry

import (
	"time"
)

// IndexEntry 首页
type IndexEntry struct {
	UserEntry              //当前用户
	TalkList   []UserEntry //会话列表
	FriendList []UserEntry //好友列表
	GroupList  interface{} //群组列表
}

// UserEntry 用户
type UserEntry struct {
	ID         string
	Account    string
	Password   string
	Email      string
	RemarkName string
	Avatar     string
}

// Message 消息的结构体
type Message struct {
	SenderID   string
	ReceiverID string
	Time       time.Time
	Content    string
}
