package main

import (
	"net/http"

	"github.com/Scfy-Code/IM/app/account/router"
	"github.com/Scfy-Code/IM/sys"
)

func main() {
	sys.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(sys.APP.StaticDir))))
	sys.Handle("/login.scfy", router.NewloginTemplate())
	sys.Handle("/login.action", router.NewLoginRouter())
	sys.Handle("/regist.scfy", router.NewRegistTemplate())
	sys.Handle("/regist.action", router.NewRegistRouter())
	sys.ListenAndServe(":8086")
}
