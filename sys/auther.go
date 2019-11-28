package sys

import "net/http"

// Authenticator 认证器接口
type Authenticator interface {
	Auth(w http.ResponseWriter, r *http.Request) bool
}
type authenticator struct {
}

func (a authenticator) Auth(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("SESSIONID")
	if err != nil {
		WarnLogger.Println(err.Error())
		return false
	}
	sessionID := cookie.Value
	WarnLogger.Println(sessionID)
	return true
}

// NewAuthenticator 创建认证器
func NewAuthenticator() Authenticator {
	return &authenticator{}
}
