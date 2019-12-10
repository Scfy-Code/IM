package router

import "net/http"
type registTemplate struct{

}
func (rt registTemplate) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

type registRouter struct {
}

// NewRegistRouter 创建注册路由
func NewRegistRouter() http.Handler {
	return registRouter{}
}
func (rr registRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
