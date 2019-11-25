package talker

import (
	"encoding/json"
	"net/http"

	"github.com/Scfy-Code/IM/app/service"
)

type deleteTalker struct {
	talkerService service.TalkerService
}

func (dt deleteTalker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var msg = map[string]interface{}{
		"status":  "success",
		"message": "已退出群聊",
	}
	if r.Method == "POST" {
		var (
			bindID = r.PostFormValue("bindID")
		)
		if dt.talkerService.DeleteTalker(bindID) {
			msg["status"] = "success"
			msg["message"] = "删除成功！"
		}
	} else {
		msg["status"] = "failure"
		msg["message"] = "删除失败！"
	}
	result, err := json.Marshal(msg)
	if err != nil {

	}
	w.Write(result)
}

// NewDeleteTalkerRouter 新建删除好友路由
func NewDeleteTalkerRouter() http.Handler {
	return deleteTalker{
		service.NewTalkerService("talkerService"),
	}
}
