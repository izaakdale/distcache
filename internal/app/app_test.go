package app_test

import (
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/izaakdale/distcache/internal/store"
)

type test struct{}

func (t *test) Ping() *redis.StatusCmd { return &redis.StatusCmd{} }
func (t *test) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return redis.NewStatusResult("set", nil)
}
func (t *test) Get(key string) *redis.StringCmd {
	return redis.NewStringResult("hello", nil)
}

func TestXxx(t *testing.T) {
	tx := &test{}
	store.WithTransactioner(tx)
}
