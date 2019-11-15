package main

import (
	"net/http"

	"github.com/Scfy-Code/IM/sys/router"
	"github.com/Scfy-Code/IM/sys/router/talker"
	"github.com/Scfy-Code/IM/sys/router/team"
	"github.com/Scfy-Code/IM/app"
)

func main() {
	http.ListenAndServe(":8088", nil)
}
func init() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(pkg.APP.StaticDir))))
	http.Handle("/", router.NewIndexTemplateRouter())
	http.Handle("/login.scfy", router.NewLoginRouter())
	http.Handle("/regist.scfy", router.NewRegistRouter())
	http.Handle("/login.action", router.NewLoginRouter())
	http.Handle("/regist.action", router.NewRegistRouter())
	http.Handle("/delete_talker.action", talker.NewDeleteTalkerRouter())
	http.Handle("/select_talkerInfo.action", talker.NewSelectTalkerRouter())
	http.Handle("/quit_team.action", team.NewQuitTeamRouter())
}
