package mock

import (
	"github.com/go-redis/redis"
	"github.com/izaakdale/distcache/internal/store"
	"time"
)

var _ store.RedisTransactioner = (*Redis)(nil)

type Redis struct {
	Cache map[string]string
}

func (r *Redis) Ping() *redis.StatusCmd {
	return redis.NewStatusResult("", nil)
}
func (r *Redis) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return redis.NewStatusResult(key, nil)
}
func (r *Redis) Get(key string) *redis.StringCmd {
	return redis.NewStringResult(Val1, nil)
}
func (r *Redis) Keys(pattern string) *redis.StringSliceCmd {
	return redis.NewStringSliceCmd(Key1, Key2)
}
func (r *Redis) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	return redis.NewScanCmdResult([]string{Key1, Key2}, 0, nil)
}
