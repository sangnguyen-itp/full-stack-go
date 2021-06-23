package caching

import (
	"full-stack-go/app"
	"github.com/go-redis/redis/v7"
	"os"
)

type RedisService struct {
	client *redis.Client
}

func (r RedisService) Set(key, value string) error {
	return r.client.Set(key, value, 0).Err()
}

func (r RedisService) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r RedisService) GetByPrefix(prefix string) ([]string, error) {
	panic("implement me")
}

func (r RedisService) Delete(key string) error {
	panic("implement me")
}

func (r RedisService) DeleteByPrefix(prefix string) error {
	panic("implement me")
}

var _ CacheServiceInterface = &RedisService{}

func NewRedisService(env *app.Environment, dbNo int) *RedisService {
	connectKey := env.UseCacheDB + "_" + env.EnvMode + "_"
	return &RedisService{
		redis.NewClient(&redis.Options{
		Addr: os.Getenv(connectKey + "HOST") + ":" + os.Getenv(connectKey + "PORT"),
		Password: os.Getenv(connectKey + "PASSWORD"),
		DB: dbNo,
	})}
}
