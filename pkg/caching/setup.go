package caching

import (
	"fmt"
	"full-stack-go/app"
	"full-stack-go/constants"
)

type CacheServiceInterface interface {
	Set(key, value string) error
	Get(key string) (string, error)
	GetByPrefix(prefix string) ([]string, error)
	Delete(key string) error
	DeleteByPrefix(prefix string) error
}

type CacheService struct {
	Auth           CacheServiceInterface
	Logging        CacheServiceInterface
	UserActivities CacheServiceInterface
}

var Cache *CacheService

func NewCacheService(env *app.Environment) {
	switch env.UseCacheDB {
	case constants.REDIS:
		Cache = &CacheService{
			Auth: NewRedisService(env, 0),
			Logging: NewRedisService(env, 1),
			UserActivities: NewRedisService(env, 2),
		}
	case constants.MEMCACHED:
		memcachedService := NewMemcachedServiceService(env)
		Cache = &CacheService{
			Auth: memcachedService,
			Logging: memcachedService,
			UserActivities: memcachedService,
		}
	default:
		panic(fmt.Errorf("caching db invalid: %s. Caching service only support for [redis, memcached]", env.UseCacheDB))
	}
}