package sys

import "net/http"

// Authenticator 认证器接口
type Authenticator interface {
	Auth(w http.ResponseWriter, r *http.Request) bool
}
type authenticator struct {
}

func (a authenticator) Auth(w http.ResponseWriter, r *http.Request) bool {
	return false
}

// NewAuthenticator 创建认证器
func NewAuthenticator() Authenticator {
	return &authenticator{}
}
