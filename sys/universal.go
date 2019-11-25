package sys

import (
	"database/sql"

	"github.com/go-redis/redis"
)

var (
	// SQLClient sql客户端
	SQLClient *sql.DB
	// RedisClient redis客户端
	RedisClient redis.UniversalClient
	// logger 日志对象
	logger = NewWarnLogger()
	// APP 配置对象
	APP application = application{}
	// temp 模板对象
	temp Template
)
