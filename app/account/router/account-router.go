package router

import (
	"net/http"

	"github.com/Scfy-Code/IM/app/account/service"
)

type loginRouter struct {
	accountService service.AccountService
}

// NewLoginRouter 创建登录请求路由
func NewLoginRouter() http.Handler {
	return loginRouter{
		service.NewAccountService(),
	}
}
func (lr loginRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var (
			email    string = r.PostFormValue("email")
			password string = r.PostFormValue("password")
		)
		lr.accountService.SelectAccount(email, password)
	}
}

type registRouter struct {
	accountService service.AccountService
}

// NewRegistRouter 创建注册路由
func NewRegistRouter() http.Handler {
	return registRouter{
		service.NewAccountService(),
	}
}
func (rr registRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var (
			email     string = r.PostFormValue("email")
			password  string = r.PostFormValue("password")
			possword0 string = r.PostFormValue("password0")
		)
		rr.accountService.InsertAccount(email, password, possword0)
	}
}
