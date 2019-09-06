package users

import (
	"html/template"
	"net/http"

	"github.com/Scfy-Code/scfy-im/entry"
)

// SignupView 注册页面处理器
var SignupView = &signupView{entry.Views["sign_up.scfy"]}

// SignupAction 注册请求处理器
var SignupAction = &signupAction{"", entry.Views["sign_up.scfy"]}

// SignUpView 注册页面结构体
type signupView struct {
	signupTemplate *template.Template
}

func (sv signupView) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sv.signupTemplate.Execute(w, nil)
}

type signupAction struct {
	redirectURL    string             //注册成功重定向的RUL
	signupTemplate *template.Template //注册失败转发的页面
}

func (sa signupAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
