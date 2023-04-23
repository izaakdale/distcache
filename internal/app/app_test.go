package app_test

import (
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/izaakdale/distcache/internal/store"
)

type test struct{}

func (t *test) Ping() *redis.StatusCmd {
	return &redis.StatusCmd{}
}
func (t *test) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return redis.NewStatusResult("set", nil)
}
func (t *test) Get(key string) *redis.StringCmd {
	return redis.NewStringResult("hello", nil)
}
func (t *test) Keys(pattern string) *redis.StringSliceCmd {
	return redis.NewStringSliceCmd("mikey", "keith")
}
func (t *test) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	return redis.NewScanCmdResult([]string{"mikey", "keith"}, 0, nil)
}

func TestXxx(t *testing.T) {
	tx := &test{}
	store.WithTransactioner(tx)
}
