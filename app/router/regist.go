package router

import "net/http"

type registRouter struct {
}

// NewRegistRouter 创建注册路由
func NewRegistRouter() http.Handler {
	return registRouter{}
}
func (rr registRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
