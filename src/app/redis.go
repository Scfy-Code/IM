package app

import cache "github.com/go-redis/redis"

var (
	// RedisClient redis客户端
	RedisClient cache.UniversalClient
)

//redis配置
type redis struct {
	Cluster        bool                  `json:"cluster"`        //是否启用redis集群
	StandOptions   *cache.Options        `json:"standOptions"`   //redis单机配置信息
	ClusterOptions *cache.ClusterOptions `json:"clusterOptions"` //redis集群配置信息
}
