package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	cache "github.com/go-redis/redis"
)

// Session session域的操作接口
type Session interface {
	GetData(field string) map[string]interface{}           //获取session域中field字段的值
	SetData(field string, value interface{}) bool          //存储和修改session域中field字段的值为value
	DelData(fields ...string) bool                         //删除session域中fields字段
	IsExist(field string) bool                             //验证session域中field字段是否存在
	SessionID() string                                     //获取session的sessionid
	RmSession(w http.ResponseWriter, r *http.Request) bool //销毁session对象关闭session域
}

type session struct {
	sessionID   string
	redisClient cache.UniversalClient
}

// NewSession 通过请求获取session域
func NewSession(w http.ResponseWriter, r *http.Request) Session {
	var sessionID string
	cookie, err := r.Cookie("SESSIONID")
	if err != nil || cookie == nil {
		sessionID = fmt.Sprintf("%d", time.Now().Unix())
		cookie0 := &http.Cookie{Name: "SESSIONID", Value: sessionID, Path: "/"}
		r.AddCookie(cookie0)
		if w != nil {
			http.SetCookie(w, cookie0)
		}
	} else {
		sessionID = cookie.Value
	}
	if w != nil {
		http.SetCookie(w, cookie)
	}
	return &session{sessionID, RedisClient}
}
func (s session) GetData(field string) map[string]interface{} {
	result0, err0 := s.redisClient.HGet(s.SessionID(), field).Result()
	if err0 != nil {
		WarnLogger.Printf("session刷新生命周期出错！错误信息：%s", err0.Error())
		return nil
	}
	_, err1 := s.redisClient.Expire(s.SessionID(), time.Minute*30).Result()
	if err1 != nil {
		WarnLogger.Printf("session刷新生命周期出错！错误信息：%s", err1.Error())
		return nil
	}
	var m = make(map[string]interface{})
	data := []byte(result0)
	err2 := json.Unmarshal(data, &m)
	if err2 != nil {
		WarnLogger.Printf("session刷新生命周期出错！错误信息：%s", err2.Error())
		return nil
	}
	return m
}
func (s session) SetData(field string, value interface{}) bool {
	var data []byte
	v, ok := value.(string)
	if ok {
		data = []byte(v)
	}
	data, err0 := json.Marshal(value)
	if err0 != nil {
		WarnLogger.Printf("序列化session数据出错！错误信息：%s", err0.Error())
		return false
	}
	result, err1 := s.redisClient.HSet(s.SessionID(), field, data).Result()
	if err1 != nil {
		WarnLogger.Printf("存储session数据出错！错误信息：%s", err1.Error())
		return false
	}
	_, err2 := s.redisClient.Expire(s.SessionID(), time.Minute*30).Result()
	if err2 != nil {
		WarnLogger.Printf("设置session声明周期出错！错误信息：%s", err2.Error())
	}
	return result
}
func (s session) DelData(fields ...string) bool {
	_, err0 := s.redisClient.HDel(s.SessionID(), fields...).Result()
	_, err1 := s.redisClient.Expire(s.SessionID(), time.Minute*30).Result()
	if err0 != nil {
		WarnLogger.Printf("删除session数据出错！错误信息：%s", err0.Error())
		return false
	}
	if err1 != nil {
		WarnLogger.Printf("设置session声明周期出错！错误信息：%s", err1.Error())
	}
	return true
}
func (s session) IsExist(field string) bool {
	result, err0 := s.redisClient.HExists(s.SessionID(), field).Result()
	if err0 != nil {
		WarnLogger.Printf("验证session数据出错！错误信息：%s", err0.Error())
		return false
	}
	_, err1 := s.redisClient.Expire(s.SessionID(), time.Minute*30).Result()
	if err1 != nil {
		WarnLogger.Printf("设置session生命周期出错！错误信息：%s", err0.Error())
	}
	return result
}
func (s session) SessionID() string {
	return s.sessionID
}
func (s session) RmSession(w http.ResponseWriter, r *http.Request) bool {
	_, err := s.redisClient.Del(s.SessionID()).Result()
	if err != nil {
		WarnLogger.Printf("移除session出错！错误信息：%s", err.Error())
		return false
	}
	cookie := &http.Cookie{Name: "SESSIONID", MaxAge: -1}
	r.AddCookie(cookie)
	http.SetCookie(w, cookie)
	return true
}
