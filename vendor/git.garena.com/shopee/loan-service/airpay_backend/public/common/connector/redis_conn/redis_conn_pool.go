package redis_conn

import (
	"strings"
	"time"

	"github.com/go-redis/redis"
)

type RedisConnConfig struct {
	RedisAddr     string
	RedisPassword string
	DB            int

	RedisPoolSize        int
	DialTimeoutMs        int
	ReadTimeoutMs        int
	WriteTimeoutMs       int
	IdleTimeoutMs        int
	IdleCheckFrequencyMs int
	ClusterMode          bool
}

func GetRedisConn(c *RedisConnConfig) (redis.UniversalClient, error) {
	redisDB := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    strings.Split(c.RedisAddr, ","),
		Password: c.RedisPassword,
		DB:       c.DB,

		PoolSize:           c.RedisPoolSize,
		DialTimeout:        time.Duration(c.DialTimeoutMs) * time.Millisecond,
		ReadTimeout:        time.Duration(c.ReadTimeoutMs) * time.Millisecond,
		WriteTimeout:       time.Duration(c.WriteTimeoutMs) * time.Millisecond,
		IdleTimeout:        time.Duration(c.IdleTimeoutMs) * time.Millisecond,
		IdleCheckFrequency: time.Duration(c.IdleCheckFrequencyMs) * time.Millisecond,
	})

	if err := redisDB.Ping().Err(); err != nil {
		return nil, err
	}
	return redisDB, nil
}