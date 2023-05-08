package store

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"log"
	"time"
)

var (
	_ Transactioner = (*Client)(nil)

	ErrRecordNotFound = errors.New("no record matches that key") //TODO add return where applicable
)

type Client struct {
	RedisTransactioner
}

type Transactioner interface {
	Insert(k, v string, ttl int) error
	Fetch(k string) (string, error)
	AllKeys(pattern string) ([]string, error)
	//Reader() io.Reader
}

type RedisTransactioner interface {
	Ping() *redis.StatusCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(key string) *redis.StringCmd
	Scan(cursor uint64, match string, count int64) *redis.ScanCmd
}

type Config struct {
	RedisAddr string
	Password  string
	DB        int
	RecordTTL int
}

func New(opts ...options) (*Client, error) {
	if opts == nil {
		return nil, errors.New("must provide a config or redis transactioner.")
	}
	var client *Client
	for _, opt := range opts {
		if opt.cfg != nil {
			client = &Client{
				redis.NewClient(&redis.Options{
					Addr:     opt.cfg.RedisAddr,
					Password: opt.cfg.Password,
					DB:       opt.cfg.DB,
				}),
			}
			if client.Ping().Err() != nil {
				log.Printf("backing off connection for 3s\n")
				time.Sleep(3 * time.Second)
				return New(opt)
			}
		}
		if opt.txer != nil {
			client = &Client{opt.txer}
		}
	}
	return client, nil
}

func (c *Client) Insert(k, v string, ttl int) error {
	t := time.Second * time.Duration(ttl)
	return c.Set(k, v, t).Err()
}
func (c *Client) Fetch(k string) (string, error) {
	val, err := c.Get(k).Result()
	if err != nil {
		if err == redis.Nil {
			return "", ErrRecordNotFound
		}
		return "", err
	}
	return val, err
}
func (c *Client) AllKeys(pattern string) ([]string, error) {
	var keys []string
	iter := c.Scan(0, pattern, 0).Iterator()
	for iter.Next() {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return keys, nil
}

func WithTransactioner(txer RedisTransactioner) options {
	return options{txer: txer}
}
func WithConfig(cfg Config) options {
	return options{cfg: &cfg}
}

type options struct {
	txer RedisTransactioner
	cfg  *Config
}
