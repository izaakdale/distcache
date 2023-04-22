package store

import (
	"time"

	"github.com/go-redis/redis"
)

var client Transactioner
var ttl time.Duration

type Transactioner interface {
	Ping() *redis.StatusCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(key string) *redis.StringCmd
}

func Init(opts ...options) error {
	for _, opt := range opts {
		if opt.txer != nil {
			client = opt.txer
		}
		if opt.cfg != nil {
			ttl = time.Second * time.Duration(opt.cfg.RecordTTL)
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

func Insert(k, v string) error {
	return client.Set(k, v, ttl).Err()
}
func Fetch(k string) (string, error) {
	return client.Get(k).Result()
}
