package main

import (
	"net/http"

	"github.com/Scfy-Code/IM/app/router"
	"github.com/Scfy-Code/IM/app/router/talker"
)

func main() {
	http.ListenAndServe(":8088", nil)
}
func init() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../web/static/"))))
	http.Handle("/index.scfy", router.NewIndexTemplateRouter())
	http.Handle("/delete_talker.action", talker.NewDeleteTalkerRouter())
	http.Handle("/talkerInfo.action", talker.NewSelectTalkerRouter())
	http.Handle("/quit_team.action", nil)
}
