package router

import (
	"encoding/json"
	"net/http"
)

type talker struct {
}

func (t talker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var msg map[string]interface{} = make(map[string]interface{}, 2)
	msg["status"] = "success"
	msg["message"] = "删除成功！"
	result, err := json.Marshal(msg)
	if err != nil {

	}
	w.Write(result)
}

// NewTalkerAction 返回聊天路由
func NewTalkerAction() http.Handler {
	return &talker{}
}
