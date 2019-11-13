package router

import (
	"net/http"

	"github.com/Scfy-Code/IM/app/service"
	_ "github.com/Scfy-Code/IM/pkg/conf" //只使用初始化方法
	"github.com/Scfy-Code/IM/pkg/view"
)

type indexTemplate struct {
	talkerService service.TalkerService
	teamService   service.TeamService
}

func (it indexTemplate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		data map[string][]map[string]interface{} = make(map[string][]map[string]interface{})
	)
	data["talkerList"] = it.talkerService.SelectTalkers("111111111")
	data["teamList"] = it.teamService.SelectTeams("111111111")
	view.ReturnTemplate("index.scfy").Execute(w, data)
}

// NewIndexTemplateRouter 返回首页模板路由
func NewIndexTemplateRouter() http.Handler {
	return indexTemplate{
		service.NewTalkerService(),
		service.NewTeamService(),
	}
}
