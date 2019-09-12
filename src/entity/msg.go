package entity

import (
	"time"
)

// MsgType 消息类型
type MsgType uint8

const (
	// FILE 文件类型
	FILE MsgType = iota
	// IMAGE 图片类型
	IMAGE MsgType = iota
	// TEXT 文本类型
	TEXT MsgType = iota
)

// Msg 消息结构体
type Msg struct {
	sender   string
	receiver string
	content  string
	msgType  MsgType
	date     time.Time
}

// NewTextMsg 创建一个文本消息对象
func NewTextMsg(sender, receiver, content string) *Msg {
	msg := new(Msg)
	msg.SetSender(sender)
	msg.SetReceiver(receiver)
	msg.SetContent(content)
	msg.SetMsgType(TEXT)
	return msg
}

// NewFiletMsg 创建一个文件消息对象
func NewFiletMsg(sender, receiver, content string) *Msg {
	msg := new(Msg)
	msg.SetSender(sender)
	msg.SetReceiver(receiver)
	msg.SetContent(content)
	msg.SetMsgType(FILE)
	return msg
}

// NewImageMsg 创建一个文本消息对象
func NewImageMsg(sender, receiver, content string) *Msg {
	msg := new(Msg)
	msg.SetSender(sender)
	msg.SetReceiver(receiver)
	msg.SetContent(content)
	msg.SetMsgType(IMAGE)
	msg.SetDate()
	return msg
}

// SetSender 设置发送者
func (m Msg) SetSender(sender string) {
	m.sender = sender
}

// SetReceiver 设置接收者
func (m Msg) SetReceiver(receiver string) {
	m.receiver = receiver
}

// SetContent 设置消息内容
func (m Msg) SetContent(content string) {
	m.content = content
}

// SetMsgType 设置消息类型
func (m Msg) SetMsgType(msgType MsgType) {
	m.msgType = msgType
}

// SetDate 设置消息发生时间
func (m Msg) SetDate() {
	m.date = time.Now()
}

// GetSender 获取发送者
func (m Msg) GetSender() string {
	return m.sender
}

// GetReceiver 获取接收者
func (m Msg) GetReceiver() string {
	return m.receiver
}

// GetContent 获取消息内容
func (m Msg) GetContent() string {
	return m.content
}

// GetMsgType 获取消息类型
func (m Msg) GetMsgType() MsgType {
	return m.msgType
}

// GetDate 获取消息的发生时间
func (m Msg) GetDate() time.Time {
	return m.date
}

// IsNotNil 消息对象的判空
func (m Msg) IsNotNil() bool {
	if m.sender == "" {
		return false
	}
	if m.receiver == "" {
		return false
	}
	if m.content == "" {
		return false
	}
	return true
}
