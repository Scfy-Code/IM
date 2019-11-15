package authen

import "net/http"

type auth struct {
	*http.ServeMux
}

func (a auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.ServeHTTP(w, r)
}

// NewAuthen 创建认证对象
func NewAuthen() http.Handler {
	return &auth{
		http.DefaultServeMux,
	}
}
