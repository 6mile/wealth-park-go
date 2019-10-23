package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	t.Run("should succeed to get config instance", func(t *testing.T) {
		c := GetInstance()
		require.True(t, c.Port > 0, "should have port > 0")
		require.Equal(t, c, GetInstance(), "should be the same instance")
		require.NotEmpty(t, c.GetAddr(), "should have address set")
	})
}
