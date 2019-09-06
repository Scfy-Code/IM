package users

import (
	"html/template"
	"net/http"

	"github.com/Scfy-Code/scfy-im/entry"
)

// SigninView 登录页面处理器
var SigninView = &signinView{entry.Views["sign_in.scfy"]}

// SigninAction 登录请求处理器
var SigninAction = &signinAction{"", entry.Views["sign_in.scfy"]}

// signinView 登录页面结构体
type signinView struct {
	signinTemplate *template.Template
}

func (sv signinView) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sv.signinTemplate.Execute(w, nil)
}

// signinAction 登录请求的结构体
type signinAction struct {
	redirectURL    string             //登录成功重定向的页面
	signinTemplate *template.Template //登陆失败转发的页面
}

// 处理登录请求：登录成功重定向至指定页面/登录失败返回登录页面
func (sa signinAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, sa.redirectURL, http.StatusFound)
}
