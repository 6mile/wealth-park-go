package core_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yashmurty/wealth-park/wpark/core"
)

func TestNewProduct(t *testing.T) {
	t.Run("should succeed and return new product resource", func(t *testing.T) {
		b, err := core.NewProduct(core.NewProductArgs{
			Name: "Test product name",
		})
		require.NoError(t, err)
		require.NotEmpty(t, b.ID)
		require.Equal(t, b.Name, "Test product name")
	})

	t.Run("should fail due to missing name", func(t *testing.T) {
		_, err := core.NewProduct(core.NewProductArgs{})
		require.Error(t, err)
	})
}
