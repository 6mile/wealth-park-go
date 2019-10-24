package mysql

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yashmurty/wealth-park/wpark/core"
)

type PurchaserProductModelTestData struct {
	model                 *PurchaserProductModel
	testPurchaserProduct1 *core.PurchaserProduct
	testPurchaserProduct2 *core.PurchaserProduct
}

func NewPurchaserProductModelTestData() *PurchaserProductModelTestData {
	t := PurchaserProductModelTestData{}
	SetupDBHandle()

	t.testPurchaserProduct1, _ = core.NewPurchaserProduct(core.NewPurchaserProductArgs{
		ID:                "PURCHASER-PRODUCT-1",
		PurchaserID:       "PURCHASER-1",
		ProductID:         "PRODUCT-1",
		PurchaseTimestamp: time.Now().Unix(),
	})

	t.testPurchaserProduct2, _ = core.NewPurchaserProduct(core.NewPurchaserProductArgs{
		ID:                "PURCHASER-PRODUCT-2",
		PurchaserID:       "PURCHASER-2",
		ProductID:         "PRODUCT-2",
		PurchaseTimestamp: time.Now().Unix(),
	})

	t.model = NewPurchaserProductModel()

	return &t
}

func TestPurchaserProductCreateTable(t *testing.T) {
	d := NewPurchaserProductModelTestData()

	t.Run("should succeed and create table", func(t *testing.T) {
		err := d.model.CreateTable(context.Background(), true)
		require.NoError(t, err)
	})
}

func TestPurchaserProductCreate(t *testing.T) {
	d := NewPurchaserProductModelTestData()

	t.Run("should succeed and create purchaser", func(t *testing.T) {
		// Create runs successfully.
		err := d.model.Create(context.Background(), d.testPurchaserProduct1)
		require.NoError(t, err)
	})

	t.Run("should fail due to duplicate primary key", func(t *testing.T) {
		// Create fails due to duplicate primary key.
		err := d.model.Create(context.Background(), d.testPurchaserProduct1)
		require.Error(t, err)
	})

	t.Run("should succeed and create purchaser", func(t *testing.T) {
		// Create runs successfully.
		err := d.model.Create(context.Background(), d.testPurchaserProduct2)
		require.NoError(t, err)
	})

	t.Run("should fail due to duplicate name", func(t *testing.T) {
		// Create fails due to duplicate name.
		err := d.model.Create(context.Background(), d.testPurchaserProduct2)
		require.Error(t, err)
	})

}
