package store_test

import (
	"testing"

	"github.com/izaakdale/distcache/internal/store"
	"github.com/izaakdale/distcache/internal/store/mock"
	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {
	tx := &mock.MockStorePass{}
	err := store.Init(store.WithTransactioner(tx))
	defer store.Reset()
	require.NoError(t, err)

	t.Run("store pass", func(t *testing.T) {
		err := store.Insert(mock.Key1, mock.Val1, 0)
		require.NoError(t, err)
	})

	t.Run("fetch pass", func(t *testing.T) {
		val, err := store.Fetch(mock.Key1)
		require.NoError(t, err)
		require.Equal(t, mock.Val1, val)
	})

	t.Run("get keys pass", func(t *testing.T) {
		keys, err := store.AllKeys("")
		require.NoError(t, err)
		require.Equal(t, mock.Key1, keys[0])
		require.Equal(t, mock.Key2, keys[1])
	})

	t.Run("ttl pass", func(t *testing.T) {
		ttl, err := store.GetTTL(mock.Key1)
		require.NoError(t, err)
		require.Equal(t, int32(-1), *ttl)
	})
}
