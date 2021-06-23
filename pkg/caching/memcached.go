package caching

import (
	"fmt"
	"full-stack-go/app"
	"github.com/bradfitz/gomemcache/memcache"
	"os"
)

type MemcachedService struct {
	client *memcache.Client
}

func (r MemcachedService) Set(key, value string) error {
	return r.client.Set(&memcache.Item{
		Key: key,
		Value: []byte(value),
	})
}

func (r MemcachedService) Get(key string) (string, error) {
	data, err := r.client.Get(key)
	if err != nil {
		return "", err
	}
	return string(data.Value), nil
}

func (r MemcachedService) GetByPrefix(prefix string) ([]string, error) {
	panic("implement me")
}

func (r MemcachedService) Delete(key string) error {
	panic("implement me")
}

func (r MemcachedService) DeleteByPrefix(prefix string) error {
	panic("implement me")
}

var _ CacheServiceInterface = &RedisService{}

func NewMemcachedServiceService(env *app.Environment) *MemcachedService {
	connectKey := env.UseCacheDB + "_" + env.EnvMode + "_"
	mem := &MemcachedService{
		client: memcache.New(
			fmt.Sprintf("%s:%s",
				os.Getenv(connectKey + "HOST"),
				os.Getenv(connectKey + "PORT"))),
	}
	// trigger connection
	if err := mem.Set("trigger","success"); err != nil {
		panic(err)
	}
	return mem
}
