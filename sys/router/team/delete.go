package team

import (
	"encoding/json"
	"net/http"

	"github.com/Scfy-Code/IM/sys/service"
)

type quitTeam struct {
	service.TeamService
}

// NewQuitTeamRouter 创建退出群聊路由器
func NewQuitTeamRouter() http.Handler {
	return quitTeam{
		service.NewTeamService("teamservice"),
	}
}
func (qt quitTeam) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var msg = map[string]interface{}{
		"status":  "success",
		"message": "已退出群聊",
	}
	if r.Method == "POST" {
		var bindID = r.PostFormValue("bindID")
		if qt.DeleteTeam(bindID) {
			msg["status"] = "success"
			msg["message"] = "已退出群聊"
		}
	}
	result, err := json.Marshal(msg)
	if err != nil {

	}
	w.Write(result)
}
