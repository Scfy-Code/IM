package sys

import (
	"net/http"
)

// UniversalServerHandler 通用处理器
type UniversalServerHandler struct {
	*http.ServeMux
	Authenticator
}

// NewUniversalServerHandler 创建通用处理器
func NewUniversalServerHandler() *UniversalServerHandler {

	return &UniversalServerHandler{
		http.NewServeMux(),
		NewAuthenticator(),
	}
}

// AuthHandle 注册需要认证的处理器
func (ush UniversalServerHandler) AuthHandle(pattern string, handler http.Handler) {
	ush.Handle(pattern, handler)
}
func (ush UniversalServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, pattern := ush.Handler(r)
	if handler != nil {
		handler, pattern = http.DefaultServeMux.Handler(r)
		if !ush.Auth(w, r) {
			logger.Println(pattern)
			return
		}
	}
	handler.ServeHTTP(w, r)

}
