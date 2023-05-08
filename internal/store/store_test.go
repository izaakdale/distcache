package store_test

import (
	"github.com/izaakdale/distcache/internal/store"
	"github.com/izaakdale/distcache/internal/store/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func initTest(t *testing.T) *store.Client {
	client, err := store.New(store.WithTransactioner(&mock.Redis{
		Cache: map[string]any{},
	}))
	require.NoError(t, err)
	require.NotNil(t, client)
	require.NotNil(t, client.RedisTransactioner)

	return client
}

func TestClient(t *testing.T) {
	client := initTest(t)

	t.Run("store and fetch", func(t *testing.T) {
		err := client.Insert(mock.Key1, mock.Val1, 0)
		require.NoError(t, err)

		val, err := client.Fetch(mock.Key1)
		require.NoError(t, err)
		require.Equal(t, mock.Val1, val)

		keys, err := client.AllKeys("")
		require.NoError(t, err)
		require.Equal(t, mock.Key1, keys[0])
		require.Equal(t, 1, len(keys))

		err = client.Insert(mock.Key2, mock.Val2, 0)
		keys, err = client.AllKeys("")
		require.NoError(t, err)
		require.Equal(t, mock.Key1, keys[0])
		require.Equal(t, mock.Key2, keys[1])
		require.Equal(t, 2, len(keys))
	})

	client = initTest(t)

	t.Run("errors", func(t *testing.T) {
		val, err := client.Fetch("non-existent-key")
		require.Empty(t, val)
		require.ErrorIs(t, err, store.ErrRecordNotFound)

		err = client.Insert(mock.Key1, mock.Val1, 0)
		require.NoError(t, err)

		val, err = client.Fetch("non-existent-key")
		require.Empty(t, val)
		require.ErrorIs(t, err, store.ErrRecordNotFound)
	})
}
