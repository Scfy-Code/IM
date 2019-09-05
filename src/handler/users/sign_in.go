package users

import (
	"net/http"
	"text/template"

	service "github.com/Scfy-Code/scfy-im/service/users"
)

// SignInView 登录页面结构体
type SignInView struct {
	ViewTemplate *template.Template
}

func (si SignInView) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	si.ViewTemplate.Execute(w, nil)
}
func NewSignInView() *SignInView {
	signInView, err := template.ParseFiles("../lib/views/users/sign_in.scfy")
	if err != nil {
		return nil
	} else {
		return &SignInView{signInView}
	}
}

// SignInAction 登录请求的结构体
type SignInAction struct {
	passPath  string             //登录成功重定向的页面
	errorView *template.Template //登陆失败转发的页面
}

func (sia SignInAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//1、获取请求参数
	err0 := r.ParseForm()
	if err0 != nil {
		sia.errorView.Execute(w, "非法数据")
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")
	//2、登录
	userService := service.NewUserService()
	user := userService.Login(email, password)
	if user == nil {
		sia.errorView.Execute(w, "用户名或密码不存在")
		return
	}
	//3、将用户信息存入session
	cookie := &http.Cookie{Name: "SESSIONID", Value: "123456789", Path: "/"}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, sia.passPath, http.StatusFound)
}

// NewSignInAction 创建一个登录请求对象
func NewSignInAction() *SignInAction {
	errorView, err := template.ParseFiles("../lib/views/users/sign_in.scfy")
	if err != nil {
		return nil
	}
	return &SignInAction{"/index.scfy", errorView}
}
