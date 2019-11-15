package router

import "net/http"

type loginRouter struct {
}

// NewLoginRouter 创建登录路由
func NewLoginRouter() http.Handler {
	return loginRouter{}
}
func (lr loginRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
