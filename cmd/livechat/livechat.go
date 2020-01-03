package main

import (
	"net/http"

	"github.com/Scfy-Code/IM/app/livechat/router"
	"github.com/Scfy-Code/IM/app/livechat/router/talker"
	"github.com/Scfy-Code/IM/app/livechat/router/team"
	"github.com/Scfy-Code/IM/sys"
)

var (
	app *http.Server
)

func init() {
	app = &http.Server{
		Addr:    ":8088",
		Handler: sys.UniversalHandler,
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../web/static"))))
}
func main() {
	sys.UniversalHandler.Handle("/", router.NewIndexTemplateRouter())
	sys.UniversalHandler.Handle("/delete_talker.action", talker.NewDeleteTalkerRouter())
	sys.UniversalHandler.Handle("/select_talkerInfo.action", talker.NewSelectTalkerRouter())
	sys.UniversalHandler.Handle("/quit_team.action", team.NewQuitTeamRouter())
	app.ListenAndServe()
}
