package redis

import (
	"fmt"

	"github.com/go-redis/redis"

	"git.garena.com/shopee/loan-service/airpay_backend/merchant/merchant_admin/global/cfg"
	"git.garena.com/shopee/loan-service/airpay_backend/public/common/connector/redis_conn"
)

var rc redis.UniversalClient

func GetRedisClient() redis.UniversalClient {
	return rc
}

func InitRedisClient() {
	redisConfig := cfg.GetGlobalConfig().Redis

	var err error

	println(fmt.Sprintf("redis[db0] init: redisConfig: %+v", redisConfig))
	rc, err = redis_conn.GetRedisConn(&redis_conn.RedisConnConfig{
		RedisAddr:            redisConfig.Host,
		RedisPassword:        redisConfig.Passwd,
		DB:                   0,
		RedisPoolSize:        redisConfig.PoolSize,
		DialTimeoutMs:        redisConfig.DialTimeoutMs,
		ReadTimeoutMs:        redisConfig.ReadTimeoutMs,
		WriteTimeoutMs:       redisConfig.WriteTimeoutMs,
		IdleTimeoutMs:        redisConfig.IdleTimeoutMs,
		IdleCheckFrequencyMs: redisConfig.IdleCheckFrequency,
		ClusterMode:          redisConfig.ClusterMode,
	})

	if err != nil {
		panic(fmt.Errorf("redis[0] init occur error. err: %+v", err))
	}

	println(fmt.Sprintf("redis[0] init: redis init successfully"))
}
