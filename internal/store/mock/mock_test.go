package mock_test

import (
	"github.com/izaakdale/distcache/internal/store"
	"github.com/izaakdale/distcache/internal/store/mock"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {
	tx := mock.New()

	t.Run("store pass", func(t *testing.T) {
		err := tx.Insert(mock.Key1, mock.Val1, 0)
		require.NoError(t, err)
	})

	t.Run("fetch pass", func(t *testing.T) {
		val, err := tx.Fetch(mock.Key1)
		require.NoError(t, err)
		require.Equal(t, mock.Val1, val)
	})

	t.Run("get keys pass", func(t *testing.T) {
		keys, err := tx.AllKeys("")
		require.NoError(t, err)
		require.Equal(t, mock.Key1, keys[0])
		require.Equal(t, mock.Key2, keys[1])
	})

	t.Run("non existent key", func(t *testing.T) {
		val, err := tx.Fetch("non-existent-key")
		require.Empty(t, val)
		require.Error(t, err)
		require.ErrorIs(t, store.ErrRecordNotFound, err)
	})
}
