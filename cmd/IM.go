package main

import (
	"net/http"

	"github.com/Scfy-Code/IM/app/router"
	"github.com/Scfy-Code/IM/app/router/talker"
	"github.com/Scfy-Code/IM/app/router/team"
	"github.com/Scfy-Code/IM/pkg"
)

func main() {
	http.ListenAndServe(":8088", nil)
}
func init() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(pkg.APP.StaticDir))))
	http.Handle("/index.scfy", router.NewIndexTemplateRouter())
	http.Handle("/delete_talker.action", talker.NewDeleteTalkerRouter())
	http.Handle("/select_talkerInfo.action", talker.NewSelectTalkerRouter())
	http.Handle("/quit_team.action", team.NewQuitTeamRouter())
}
