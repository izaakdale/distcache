package store

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

var (
	client                  Transactioner
	ErrClientNotInitialised = errors.New("store client not set")
)

type Transactioner interface {
	Ping() *redis.StatusCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(key string) *redis.StringCmd
	Scan(cursor uint64, match string, count int64) *redis.ScanCmd
	TTL(key string) *redis.DurationCmd
}

func Init(opts ...options) error {
	if opts == nil {
		return ErrClientNotInitialised
	}
	for _, opt := range opts {
		if opt.txer != nil {
			client = opt.txer
		}
		if opt.cfg != nil {
			// set up client normal redis way
			client = redis.NewClient(&redis.Options{
				Addr:     opt.cfg.RedisAddr,
				Password: opt.cfg.Password,
				DB:       opt.cfg.DB,
			})
		}
	}
	return client.Ping().Err()
}

func Reset() {
	client = nil
}

type Config struct {
	RedisAddr string
	Password  string
	DB        int
	RecordTTL int
}

func WithTransactioner(txer Transactioner) options {
	return options{txer: txer}
}
func WithConfig(cfg Config) options {
	return options{cfg: &cfg}
}

type options struct {
	txer Transactioner
	cfg  *Config
}

func Insert(k, v string, ttl int) error {
	if client == nil {
		return ErrClientNotInitialised
	}
	t := time.Second * time.Duration(ttl)
	return client.Set(k, v, t).Err()
}
func Fetch(k string) (string, error) {
	return client.Get(k).Result()
}
func AllKeys(pattern string) ([]string, error) {
	var keys []string
	iter := client.Scan(0, pattern, 0).Iterator()
	for iter.Next() {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return keys, nil
}
func GetTTL(key string) (*int32, error) {
	dur, err := client.TTL(key).Result()
	if err != nil {
		return nil, err
	}
	ttl := int32(dur.Seconds())
	return &ttl, nil
}
