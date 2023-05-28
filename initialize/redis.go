package initialize

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"simple-mall/global"
)

// InitRedis 初始化redis
func InitRedis() {
	// 建立redis连接
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.RedisConfig.Host, global.RedisConfig.Port),
		Password: global.RedisConfig.Password,
		Username: global.RedisConfig.Username,
		DB:       global.RedisConfig.DB,
	})
}
