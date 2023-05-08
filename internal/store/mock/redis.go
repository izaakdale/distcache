package mock

import (
	"github.com/go-redis/redis"
	"github.com/izaakdale/distcache/internal/store"
	"time"
)

var _ store.RedisTransactioner = (*Redis)(nil)

type Redis struct {
	Cache map[string]any
}

func Clear(r *Redis) {
	r.Cache = map[string]any{}
}

func (r *Redis) Ping() *redis.StatusCmd {
	return redis.NewStatusResult("", nil)
}
func (r *Redis) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	r.Cache[key] = value
	return redis.NewStatusResult(key, nil)
}
func (r *Redis) Get(key string) *redis.StringCmd {
	val, ok := r.Cache[key]
	if !ok {
		return redis.NewStringResult("", redis.Nil)
	}
	strval, ok := val.(string)
	if !ok {
		panic("something other than a string was stored in mock")
	}
	return redis.NewStringResult(strval, nil)
}
func (r *Redis) Keys(pattern string) *redis.StringSliceCmd {
	var keys []string
	for k, _ := range r.Cache {
		keys = append(keys, k)
	}
	return redis.NewStringSliceCmd(keys)
}
func (r *Redis) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	var keys []string
	for k, _ := range r.Cache {
		keys = append(keys, k)
	}
	return redis.NewScanCmdResult(keys, 0, nil)
}
