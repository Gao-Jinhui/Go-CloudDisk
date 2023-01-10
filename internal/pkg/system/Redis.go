package system

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var Redis *redis.Client

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	err := Redis.Ping().Err()
	if err != nil {
		Logger.Error("Redis initialization failed")
	} else {
		Logger.Info("Redis initialization succeeded")
	}
}
