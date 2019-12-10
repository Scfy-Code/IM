package talker

import (
	"net/http"

	"github.com/Scfy-Code/IM/app/livechat/service"
)

type selectTalkerList struct {
	talkerService service.TalkerService
}

func (stl selectTalkerList) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

// NewSelectTalkerRouter 查询好友路由
func NewSelectTalkerRouter() http.Handler {
	return selectTalkerList{
		service.NewTalkerService("talkerService"),
	}
}
