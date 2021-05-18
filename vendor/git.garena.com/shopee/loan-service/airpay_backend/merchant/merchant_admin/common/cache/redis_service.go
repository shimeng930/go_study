package cache

import (
	"time"

	"github.com/go-redis/redis"

	global "git.garena.com/shopee/loan-service/airpay_backend/merchant/merchant_admin/global/redis"
)

type RedisImpl struct {
}

func GetCacheImpl() *RedisImpl {
	return &RedisImpl{}
}

func (r *RedisImpl) GetCacheFromRedis(key string) (*string, error) {
	rc := global.GetRedisClient()
	result, err := rc.Get(key).Result()
	switch err {
	case nil: // got
	case redis.Nil:
		return nil, nil // key not found
	default:
		return nil, err // error occure
	}

	return &result, nil
}

func (r *RedisImpl) SetCacheIntoRedis(key string, value string, expire_seconds uint32) error {
	rc := global.GetRedisClient()
	err := rc.Set(key, value, time.Duration(expire_seconds)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisImpl) SetCacheIntoRedisNX(key string, value string, expire_seconds uint32) (bool, error) {
	rc := global.GetRedisClient()
	result, err := rc.SetNX(key, value, time.Duration(expire_seconds)*time.Second).Result()
	switch err {
	case nil:
	default:
		return false, err
	}

	return result, nil
}

func (r *RedisImpl) DeleteCacheFromRedis(key []string) (int64, error) {
	rc := global.GetRedisClient()
	result, err := rc.Del(key...).Result()
	switch err {
	case nil:
	default:
		return 0, err
	}

	return result, nil
}

func (r *RedisImpl) PushCacheIntoRedis(key string, values ...string) (int64, error) {
	rc := global.GetRedisClient()
	result, err := rc.LPush(key, values).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (r *RedisImpl) PopCacheFromRedis(key string) (*string, error) {
	rc := global.GetRedisClient()
	result, err := rc.LPop(key).Result()
	switch err {
	case nil: // got
	case redis.Nil:
		return nil, nil // key not found
	default:
		return nil, err // error occure
	}

	return &result, nil
}


// 发布
func (r *RedisImpl) PublishIntoRedis(channelName string, message string) error {
	rc := global.GetRedisClient()
	err := rc.Publish(channelName, message).Err()
	if err != nil {
		return err
	}
	return nil
}