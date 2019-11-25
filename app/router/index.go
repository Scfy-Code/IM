package router

import (
	"net/http"

	"github.com/Scfy-Code/IM/app/service"
	"github.com/Scfy-Code/IM/sys"
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
	sys.ReturnTemplate("index.scfy").Execute(w, data)
}

// NewIndexTemplateRouter 创建首页模板路由
func NewIndexTemplateRouter() http.Handler {
	return indexTemplate{
		service.NewTalkerService("talkerService"),
		service.NewTeamService("teamService"),
	}
}
