package core_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yashmurty/wealth-park/wpark/core"
)

func TestNewPurchaserProduct(t *testing.T) {
	t.Run("should succeed and return new purchaser_product resource", func(t *testing.T) {
		b, err := core.NewPurchaserProduct(core.NewPurchaserProductArgs{
			PurchaserID:       "PURCHASER-1",
			ProductID:         "PRODUCT-1",
			PurchaseTimestamp: time.Now().Unix(),
		})
		require.NoError(t, err)
		require.NotEmpty(t, b.ID)
		require.Equal(t, b.PurchaserID, "PURCHASER-1")
	})

	t.Run("should fail due to missing name", func(t *testing.T) {
		_, err := core.NewPurchaserProduct(core.NewPurchaserProductArgs{})
		require.Error(t, err)
	})
}
