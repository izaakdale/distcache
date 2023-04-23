package mock

import (
	"time"

	"github.com/go-redis/redis"
)

var (
	Key1 = "mikey"
	Val1 = "valerie"
	Key2 = "keith"
	Val2 = "vera"
)

type MockStorePass struct{}

func (t *MockStorePass) Ping() *redis.StatusCmd {
	return redis.NewStatusResult("", nil)
}
func (t *MockStorePass) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return redis.NewStatusResult(key, nil)
}
func (t *MockStorePass) Get(key string) *redis.StringCmd {
	return redis.NewStringResult(Val1, nil)
}
func (t *MockStorePass) Keys(pattern string) *redis.StringSliceCmd {
	return redis.NewStringSliceCmd(Key1, Key2)
}
func (t *MockStorePass) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	return redis.NewScanCmdResult([]string{Key1, Key2}, 0, nil)
}
func (t *MockStorePass) TTL(key string) *redis.DurationCmd {
	return redis.NewDurationResult(-1*time.Second, nil)
}
