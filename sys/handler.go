package sys

import (
	"net/http"
)

// universalServerHandler 通用处理器
type universalServerHandler struct {
	*http.ServeMux
	Authenticator
}

// newUniversalServerHandler 创建通用处理器
func newUniversalServerHandler() *universalServerHandler {
	return &universalServerHandler{
		http.NewServeMux(),
		NewAuthenticator(),
	}
}

// AuthHandle 注册需要认证的处理器
func (ush universalServerHandler) AuthHandle(pattern string, handler http.Handler) {
	ush.Handle(pattern, handler)
}
func (ush universalServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, _ := ush.ServeMux.Handler(r)
	if handler != nil {
		if !ush.Auth(w, r) {
			http.Redirect(w, r, "/login.scfy", http.StatusFound)
			return
		}
	} else {
		http.DefaultServeMux.ServeHTTP(w, r)
	}
}
