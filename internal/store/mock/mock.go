package mock

import (
	"github.com/izaakdale/distcache/internal/store"
)

var (
	Key1 = "mikey"
	Val1 = "valerie"
	Key2 = "keith"
	Val2 = "vera"
)

var (
	_ store.Transactioner = (*Store)(nil)
)

func New() *Store {
	return &Store{
		Cache: map[string]string{
			Key1: Val1,
			Key2: Val2,
		},
	}
}

type Store struct {
	Cache map[string]string
}

func (s *Store) Insert(k, v string, ttl int) error {
	s.Cache[k] = v
	return nil
}

func (s *Store) Fetch(k string) (string, error) {
	val, ok := s.Cache[k]
	if !ok {
		return "", store.ErrRecordNotFound
	}
	return val, nil
}

func (s *Store) AllKeys(pattern string) ([]string, error) {
	var ret []string
	for k, _ := range s.Cache {
		ret = append(ret, k)
	}
	return ret, nil
}
