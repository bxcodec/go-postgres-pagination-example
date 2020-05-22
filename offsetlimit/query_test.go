package offsetlimit

import (
	"context"
	"os"
	"testing"

	pagination "github.com/bxcodec/go-postgres-pagination-example"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkFetchQuery(b *testing.B) {
	// change this based on the database host
	err := os.Setenv("POSTGRES_HOST", "localhost")
	require.NoError(b, err)
	err = os.Setenv("POSTGRES_PORT", "5432")
	require.NoError(b, err)
	err = os.Setenv("POSTGRES_USER", "user")
	require.NoError(b, err)
	err = os.Setenv("POSTGRES_PASSWORD", "password")
	require.NoError(b, err)
	err = os.Setenv("POSTGRES_DATABASE", "payment")
	require.NoError(b, err)
	err = os.Setenv("DB_MAX_CONN_LIFE_TIME_S", "300")
	require.NoError(b, err)
	err = os.Setenv("DB_MAX_OPEN_CONNECTION", "100")
	require.NoError(b, err)
	err = os.Setenv("DB_MAX_IDLE_CONNECTION", "10")
	require.NoError(b, err)

	db := pagination.InitDB()
	params := FetchParam{
		Limit: 10,
	}

	// set to 10K
	// total rows = 100K, each call will fetch 10 rows
	// 100K/10 = 10K
	b.N = 10000
	for i := 0; i < b.N; i++ {
		res, nextOffset, err := FetchPayment(context.Background(), db, params)
		require.NoError(b, err)
		if nextOffset != 0 {
			params.OffSet = uint64(nextOffset)
		}
		assert.NotZero(b, len(res))
	}
}
