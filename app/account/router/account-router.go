package router

import (
	"net/http"
	"regexp"

	"github.com/Scfy-Code/IM/app/account/service"
	"github.com/Scfy-Code/IM/sys"
)

const (
	emailReg string = "^[a-zA-Z]{1}[a-zA-Z0-9]{9,17}@[(163.com)|(gmail.com)|(qq.com)]$"
	password string = "^[a-zA-Z]{1}[a-zA-Z0-9]{7,15}$"
)

type loginTemplate struct {
}

// NewloginTemplate 创建登录页面路由
func NewloginTemplate() http.Handler {
	return loginTemplate{}
}
func (lr loginTemplate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var regs map[string]string = map[string]string{"email": emailReg, "password": password, "action": "/login.action"}
		sys.ReturnTemplate("login.scfy").Execute(w, regs)
	}
}

type loginRouter struct {
	accountService service.AccountService
}

// NewLoginRouter 创建登录请求路由
func NewLoginRouter() http.Handler {
	return loginRouter{}
}
func (lr loginRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var email string = r.PostFormValue("email")
		var password string = r.PostFormValue("password")
		flag0, err0 := regexp.MatchString(emailReg, email)
		flag1, err1 := regexp.MatchString(password, password)
		if err0 != nil || err1 != nil {
			sys.ReturnTemplate("login.scfy").Execute(w, nil)
			return
		}
		if flag0 && flag1 {
			lr.accountService.SelectAccount(email, password)
		}
	}
}

type registTemplate struct {
}

func (rt registTemplate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		sys.ReturnTemplate("regist.scfy").Execute(w, nil)
	}
}

type registRouter struct {
}

// NewRegistRouter 创建注册路由
func NewRegistRouter() http.Handler {
	return registRouter{}
}
func (rr registRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
