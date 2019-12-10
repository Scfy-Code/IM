package main

import (
	"net/http"

	"github.com/Scfy-Code/IM/app/livechat/router"
	"github.com/Scfy-Code/IM/app/livechat/router/talker"
	"github.com/Scfy-Code/IM/app/livechat/router/team"
	"github.com/Scfy-Code/IM/sys"
)

func main() {
	sys.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(sys.APP.StaticDir))))
	sys.AuthHandle("/", router.NewIndexTemplateRouter())
	sys.AuthHandle("/delete_talker.action", talker.NewDeleteTalkerRouter())
	sys.AuthHandle("/select_talkerInfo.action", talker.NewSelectTalkerRouter())
	sys.AuthHandle("/quit_team.action", team.NewQuitTeamRouter())
	sys.ListenAndServe(":8088")
}
