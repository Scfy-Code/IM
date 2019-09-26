package main

import (
	"net/http"

	"github.com/Scfy-Code/IM/app/router"
)

func main() {
	http.ListenAndServe(":8088", nil)
}
func init() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../web/static/"))))
	http.Handle("/index.scfy", router.NewIndexTemplateRouter())
}
