package mysql

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yashmurty/wealth-park/wpark/core"
)

type PurchaserModelTestData struct {
	model          *PurchaserModel
	testPurchaser1 *core.Purchaser
	testPurchaser2 *core.Purchaser
}

func NewPurchaserModelTestData() *PurchaserModelTestData {
	t := PurchaserModelTestData{}
	SetupDBHandle()

	t.testPurchaser1, _ = core.NewPurchaser(core.NewPurchaserArgs{
		ID:   "PRODUCT-1",
		Name: "Test purchaser 1 name",
	})

	t.testPurchaser2, _ = core.NewPurchaser(core.NewPurchaserArgs{
		ID:   "PRODUCT-2",
		Name: "Test purchaser 2 name",
	})

	t.model = NewPurchaserModel()

	return &t
}

func TestPurchaserCreateTable(t *testing.T) {
	d := NewPurchaserModelTestData()

	t.Run("should succeed and create table", func(t *testing.T) {
		err := d.model.CreateTable(context.Background(), true)
		require.NoError(t, err)
	})
}

func TestPurchaserCreate(t *testing.T) {
	d := NewPurchaserModelTestData()

	t.Run("should succeed and create purchaser", func(t *testing.T) {
		// Create runs successfully.
		err := d.model.Create(context.Background(), d.testPurchaser1)
		require.NoError(t, err)
	})

	t.Run("should fail due to duplicate primary key", func(t *testing.T) {
		// Create fails due to duplicate primary key.
		err := d.model.Create(context.Background(), d.testPurchaser1)
		require.Error(t, err)
	})

	t.Run("should succeed and create purchaser", func(t *testing.T) {
		// Create runs successfully.
		err := d.model.Create(context.Background(), d.testPurchaser2)
		require.NoError(t, err)
	})

	t.Run("should fail due to duplicate name", func(t *testing.T) {
		// Create fails due to duplicate name.
		err := d.model.Create(context.Background(), d.testPurchaser2)
		require.Error(t, err)
	})

}
