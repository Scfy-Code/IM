package main

import (
	"net/http"

	"github.com/Scfy-Code/IM/app/router"
	"github.com/Scfy-Code/IM/app/router/talker"
	"github.com/Scfy-Code/IM/app/router/team"
	"github.com/Scfy-Code/IM/sys"
)

func main() {
	sys.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(sys.APP.StaticDir))))
	sys.AuthHandle("/", router.NewIndexTemplateRouter())
	sys.Handle("/login.scfy", router.NewloginTemplate())
	sys.Handle("/login.action", router.NewLoginRouter())
	sys.Handle("/regist.scfy", router.NewRegistRouter())
	sys.Handle("/regist.action", router.NewRegistRouter())
	sys.AuthHandle("/delete_talker.action", talker.NewDeleteTalkerRouter())
	sys.AuthHandle("/select_talkerInfo.action", talker.NewSelectTalkerRouter())
	sys.AuthHandle("/quit_team.action", team.NewQuitTeamRouter())
	sys.ListenAndServe(":8088")
}
