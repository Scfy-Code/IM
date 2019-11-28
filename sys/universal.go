package sys

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

var (
	// SQLClient sql客户端
	SQLClient *sql.DB
	// RedisClient redis客户端
	RedisClient redis.UniversalClient
	// InfoLogger 常规日志对象
	InfoLogger log.Logger
	// WarnLogger 常规日志对象
	WarnLogger log.Logger
	// APP 配置对象
	APP              application = application{}
	temp             Template
	universalHandler = newUniversalServerHandler()
)

// Handle 注册路由
func Handle(pattern string, handler http.Handler) {
	http.Handle(pattern, handler)
}

// AuthHandle 注册需验证的路由
func AuthHandle(pattern string, handler http.Handler) {
	universalHandler.AuthHandle(pattern, handler)
}

// ListenAndServe 端口监听
func ListenAndServe() {
	http.ListenAndServe(":8088", universalHandler)
}
