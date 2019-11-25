package main

import "fmt"

func main() {
	add(12, 13)
	fmt.Println("输出")
	// universalHandler := app.NewUniversalHandler()
	// universalHandler.Handle("", nil)
	// universalHandler.AuthHandle("", nil)
	// server := &http.Server{
	// 	Addr:     ":8088",
	// 	ErrorLog: mylog.NewWarnLogger(),
	// 	Handler:  universalHandler,
	// }
	// server.ListenAndServe()
}
func add(a, b int) {
	return
	fmt.Println(a + b)
}
func init() {
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(app.APP.StaticDir))))
	// auther.Handle("/", router.NewIndexTemplateRouter(), authen.AuthMethod)
	// http.Handle("/login.scfy", router.NewLoginRouter())
	// http.Handle("/regist.scfy", router.NewRegistRouter())
	// http.Handle("/login.action", router.NewLoginRouter())
	// http.Handle("/regist.action", router.NewRegistRouter())
	// auther.Handle("/delete_talker.action", talker.NewDeleteTalkerRouter(), authen.AuthMethod)
	// auther.Handle("/select_talkerInfo.action", talker.NewSelectTalkerRouter(), authen.AuthMethod)
	// auther.Handle("/quit_team.action", team.NewQuitTeamRouter(), authen.AuthMethod)
}
