package mysql

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPingServer(t *testing.T) {

	t.Run("should fail since connection error simulation is true", func(t *testing.T) {
		// Simulate connection failure.
		require.Panics(t, func() {
			clientErrorMode = true
			SetupDBHandle()
		})
	})

	t.Run("should succeed and ping mysql server", func(t *testing.T) {
		clientErrorMode = false
		SetupDBHandle()

		err := PingServer(context.Background())
		require.NoError(t, err)
	})

}
