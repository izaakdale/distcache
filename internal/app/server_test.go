package app_test

import (
	"context"
	"testing"

	v1 "github.com/izaakdale/distcache/api/v1"
	"github.com/izaakdale/distcache/internal/app"
	"github.com/izaakdale/distcache/internal/store"
	"github.com/izaakdale/distcache/internal/store/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestServer(t *testing.T) {
	ctx := context.Background()
	err := store.Init(store.WithTransactioner(&mock.MockStorePass{}))
	defer store.Reset()
	require.NoError(t, err)
	s := app.Server{}

	t.Run("store pass", func(t *testing.T) {
		resp, err := s.Store(ctx, &v1.StoreRequest{
			Record: &v1.KVRecord{
				Key:   mock.Key1,
				Value: mock.Val1,
			},
			Ttl: 0,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
	})

	t.Run("fetch pass", func(t *testing.T) {
		resp, err := s.Fetch(ctx, &v1.FetchRequest{
			Key: mock.Key1,
		})
		require.NoError(t, err)
		require.Equal(t, mock.Val1, resp.Value)
	})

	t.Run("all keys pass", func(t *testing.T) {
		resp, err := s.AllKeys(ctx, &v1.AllKeysRequest{
			Pattern: "",
		})
		require.NoError(t, err)
		require.Equal(t, mock.Key1, resp.Keys[0])
		require.Equal(t, mock.Key2, resp.Keys[1])
	})

	t.Run("", func(t *testing.T) {
		m := mockstream{}
		err := s.AllRecords(&v1.AllRecordsRequest{
			Keys: []string{mock.Key1, mock.Key1},
		}, &m)
		require.NoError(t, err)

		// check twice since the mock fetches the same record
		require.Equal(t, &v1.KVRecord{Key: mock.Key1, Value: mock.Val1}, &m.records[0])
		require.Equal(t, &v1.KVRecord{Key: mock.Key1, Value: mock.Val1}, &m.records[1])
	})
}

type mockstream struct {
	records []v1.KVRecord
	grpc.ServerStream
}

func (m *mockstream) Send(rec *v1.AllRecordsResponse) error {
	m.records = append(m.records, *rec.GetRecord())
	return nil
}
