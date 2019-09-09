package users

import (
	"html/template"
	"net/http"
	"regexp"

	"github.com/Scfy-Code/scfy-im/service/users"

	scope "github.com/Scfy-Code/scfy-im/scope/imp"

	"github.com/Scfy-Code/scfy-im/entry"
)

// signinView 登录页面结构体
type signinView struct {
	redirectURL    string             //已登录用户的重定向地址
	signinTemplate *template.Template //登录页面模板
}

// NewSigninView 创建一个登录页面处理器
func NewSigninView() http.Handler {
	return &signinView{
		"index.scfy",
		entry.Views["sign_in.scfy"],
	}
}
func (sv signinView) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := scope.NewSession(r)
	if !session.IsExist("USER") {
		sv.signinTemplate.Execute(w, nil)
		return
	}
	http.Redirect(w, r, sv.redirectURL, http.StatusFound)
}

// signinAction 登录请求的结构体
type signinAction struct {
	redirectURL    string             //登录成功重定向的页面
	signinTemplate *template.Template //登陆失败转发的页面
}

// NewSigninAction 创建一个登录请求处理器
func NewSigninAction() http.Handler {
	return &signinAction{
		"/index.scfy",
		entry.Views["sign_in.scfy"],
	}
}

// 处理登录请求：登录成功重定向至指定页面/登录失败返回登录页面
func (sa signinAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := scope.NewSession(r)
	if session.IsExist("USER") {
		http.Redirect(w, r, sa.redirectURL, http.StatusFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		sa.signinTemplate.Execute(w, "登陆失败请重试")
		return
	}
	email := r.FormValue("email")
	if len(email) > 64 {
		sa.signinTemplate.Execute(w, "非法邮箱")
		return
	}

	if check, err := regexp.MatchString("^([a-z|A-Z])+([a-z|A-Z|0-9]){4,58}@((gmail)|(qq)|(163)|(126)){1}((.com)|(.cn)|(.net)){1}$", email); err != nil || !check {
		sa.signinTemplate.Execute(w, "非法邮箱地址")
		return
	}
	password := r.FormValue("password")
	if check, err := regexp.MatchString("pattern", password); err != nil || !check {
		sa.signinTemplate.Execute(w, "邮箱或密码错误！")
		return
	}
	user := users.UserService.Login(email, password)
	if !session.SetData("USER", user) {
		sa.signinTemplate.Execute(w, "登陆失败！邮箱或密码错误！")
		return
	}
	http.Redirect(w, r, sa.redirectURL, http.StatusFound)
}

// SignUpView 注册页面结构体
type signupView struct {
	redirectURL    string
	signupTemplate *template.Template
}

// NewSignupView 创建一个注册页面处理器
func NewSignupView() http.Handler {
	return &signupView{
		"/index.scfy",
		entry.Views["sign_up.scfy"],
	}
}
func (sv signupView) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sv.signupTemplate.Execute(w, nil)
}

type signupAction struct {
	redirectURL    string             //注册成功重定向的RUL
	signupTemplate *template.Template //注册失败转发的页面
}

// NewSignupAction 创建一个注册请求处理器
func NewSignupAction() http.Handler {
	return &signupAction{
		"/index.scfy",
		entry.Views["sign_up.scfy"],
	}
}
func (sa signupAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
