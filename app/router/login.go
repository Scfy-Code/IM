package router

import (
	"net/http"

	"github.com/Scfy-Code/IM/sys"
)

type loginTemplate struct {
}

// NewloginTemplate 创建登录页面路由
func NewloginTemplate() http.Handler {
	return loginRouter{}
}
func (lr loginTemplate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		sys.ReturnTemplate("login.scfy").Execute(w, nil)
	}
}

type loginRouter struct {
}

// NewLoginRouter 创建登录请求路由
func NewLoginRouter() http.Handler {
	return loginRouter{}
}
func (lr loginRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	sys.ReturnTemplate("login.scfy").Execute(w, nil)
}
