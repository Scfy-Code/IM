package users

import (
	"net/http"
	"text/template"
)

// SignUpView 注册页面结构体
type SignUpView struct {
	signUpView *template.Template
}

func (suv SignUpView) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	suv.signUpView.Execute(w, nil)
}

// NewSignUpView
func NewSignUpView() *SignUpView {
	signUpView, err := template.ParseFiles("../lib/views/users/sign_up.scfy")
	if err == nil {
		return &SignUpView{signUpView}
	} else {
		return nil
	}
}

type SignUpAction struct {
	passPath  string             //注册成功重定向的页面
	errorView *template.Template //注册失败转发的页面
}

func (sua SignUpAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//1、登录
	//2、
	cookie := &http.Cookie{Name: "SESSIONID", Value: "SESSIONID", Path: "/"}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}
func NewSignUpAction() *SignUpAction {
	errorView, err := template.ParseFiles("../lib/views/users/sign_up.scfy")
	if err != nil {
		return nil
	}
	return &SignUpAction{"/index.scfy", errorView}
}
