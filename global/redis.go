package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"time"
)

var RedisClient *redis.Client

func initRedis() {
	// 创建一个新的客户端实例
	RedisClient = redis.NewClient(&redis.Options{
		Addr:        viper.GetString("redis.addr"),
		Password:    viper.GetString("redis.password"),
		DB:          viper.GetInt("redis.db"),
		PoolSize:    viper.GetInt("redis.pool_size"),
		PoolTimeout: viper.GetDuration("redis.pool_timeout") * time.Second,
		MaxRetries:  viper.GetInt("redis.pool_max_retries"),
	})
}
