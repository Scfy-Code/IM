package sys

import (
	"net/http"
)

// autherHandler 通用处理器
type autherHandler struct {
	authenList map[string]func(http.ResponseWriter, *http.Request)
}

// newautherHandler 创建通用处理器
func newautherHandler() *autherHandler {
	return &autherHandler{
		make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

// authHandle 注册需要认证的处理器
func (ush autherHandler) authHandle(pattern string, handler http.Handler) {
	http.Handle(pattern, handler)
	ush.authenList[pattern] = handler.ServeHTTP
}

// handle 注册需要认证的处理器
func (ush autherHandler) handle(pattern string, handler http.Handler) {
	http.Handle(pattern, handler)
}

func (ush autherHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, pattern := http.DefaultServeMux.Handler(r)
	if handleFunc, ok := ush.authenList[pattern]; ok {
		handleFunc(w, r)
	} else {
		handler.ServeHTTP(w, r)
	}
}

// Handle 注册路由
func Handle(pattern string, handler http.Handler) {
	universalHandler.handle(pattern, handler)
}

// AuthHandle 注册需验证的路由
func AuthHandle(pattern string, handler http.Handler) {
	universalHandler.authHandle(pattern, handler)
}

// ListenAndServe 端口监听
func ListenAndServe(port string) {
	http.ListenAndServe(port, universalHandler)
}
