package scope

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Scfy-Code/scfy-im/logger"

	"github.com/Scfy-Code/scfy-im/database"

	"github.com/go-redis/redis"
)

// Session session域的操作接口
type Session interface {
	GetData(field string) map[string]interface{}  //获取session域中field字段的值
	SetData(field string, value interface{}) bool //存储和修改session域中field字段的值为value
	DelData(fields ...string) bool                //删除session域中fields字段
	IsExist(field string) bool                    //验证session域中field字段是否存在
	SessionID() string                            //获取session的sessionid
	RmSession() bool                              //销毁session对象关闭session域
}

// Application application域的操作接口
type Application interface {
	AddData(value ...interface{}) bool //存储和修改数据
	RemData(value ...interface{}) bool //删除application域中value元素
	IsExist(value interface{}) bool    //验证application域中是否存在value元素
}
type application struct {
	key         string                //域名称
	redisClient redis.UniversalClient //操作域的客户端对象
}

// NewApplication 根绝请求对象获取application对象
func NewApplication(key string) Application {
	return &application{key, database.RedisClient}
}
func (a application) AddData(value ...interface{}) bool {
	_, err := a.redisClient.SAdd(a.key, value...).Result()
	if err != nil {
		return false
	}
	return true
}
func (a application) RemData(value ...interface{}) bool {
	_, err := a.redisClient.SRem(a.key, value...).Result()
	if err != nil {
		return false
	}
	return true

}
func (a application) IsExist(value interface{}) bool {
	result, err := a.redisClient.SIsMember(a.key, value).Result()
	if err != nil {
		return false
	}
	return result
}

type session struct {
	sessionID   string
	redisClient redis.UniversalClient
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
	return &session{sessionID, database.RedisClient}
}
func (s session) GetData(field string) map[string]interface{} {
	result, err := s.redisClient.HGet(s.SessionID(), field).Result()
	if err != nil {
		logger.WarnPrintf("反序列化session数据报错！错误信息：%s", err.Error())
		return nil
	}
	data, err0 := json.Marshal(result)
	if err0 != nil {
		logger.WarnPrintf("反序列化session数据报错！错误信息：%s", err0.Error())
		return nil
	}
	r := make(map[string]interface{})
	err1 := json.Unmarshal(data, &r)
	if err1 != nil {
		logger.WarnPrintf("反序列化session数据报错！错误信息：%s", err1.Error())
		return nil
	}
	return r
}
func (s session) SetData(field string, value interface{}) bool {
	data, err := json.Marshal(value)
	if err != nil {
		logger.WarnPrintf("序列化session数据出错！错误信息：%s", err.Error())
		return false
	}
	var str string
	err0 := json.Unmarshal(data, &str)
	if err0 != nil {
		logger.WarnPrintf("序列化session数据出错！错误信息：%s", err0.Error())
		return false
	}
	result, err1 := s.redisClient.HSet(s.SessionID(), field, str).Result()
	logger.WarnPrintf("存储session数据出错！错误信息：%s", err1.Error())
	if err1 != nil {
		return false
	}
	return result
}
func (s session) DelData(fields ...string) bool {
	_, err := s.redisClient.HDel(s.SessionID(), fields...).Result()
	if err != nil {
		logger.WarnPrintf("删除session数据出错！错误信息：%s", err.Error())
		return false
	}
	return true
}
func (s session) IsExist(field string) bool {
	result, err := s.redisClient.HExists(s.SessionID(), field).Result()
	if err != nil {
		logger.WarnPrintf("验证session数据出错！错误信息：%s", err.Error())
		return false
	}
	return result
}
func (s session) SessionID() string {
	return s.sessionID
}
func (s session) RmSession() bool {
	_, err := s.redisClient.Del(s.SessionID()).Result()
	if err != nil {
		logger.WarnPrintf("移除session出错！错误信息：%s", err.Error())
		return false
	}
	return true
}
