package service

import (
	"context"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yashmurty/wealth-park/wpark/core"
	"github.com/yashmurty/wealth-park/wpark/mock"
)

type PurchaserProductServiceTestData struct {
	svc               *PurchaserProductService
	model             *mock.PurchaserProductModel
	PurchaserProduct1 *core.PurchaserProduct
}

func NewPurchaserProductServiceTestData() *PurchaserProductServiceTestData {
	t := PurchaserProductServiceTestData{}

	t.PurchaserProduct1, _ = core.NewPurchaserProduct(core.NewPurchaserProductArgs{
		ID:                "PRODUCT-1",
		PurchaserID:       "PURCHASER-1",
		ProductID:         "PRODUCT-1",
		PurchaseTimestamp: time.Now().Unix(),
	})

	t.svc = &PurchaserProductService{}

	t.model = &mock.PurchaserProductModel{}
	t.svc.SetPurchaserProductModel(t.model)

	return &t
}

func TestCreatePurchaserProduct(t *testing.T) {
	d := NewPurchaserProductServiceTestData()

	t.Run("should succeed and create purchase", func(t *testing.T) {
		// Mocked model function runs successfully.
		d.model.CreateFn = func(ctx context.Context, d *core.PurchaserProduct) error { return nil }

		err := d.svc.CreatePurchaserProduct(context.Background(), d.PurchaserProduct1)
		require.NoError(t, err)
		require.Equal(t, 1, d.model.CreateFnCalled)
	})

	t.Run("should fail since model function returns error", func(t *testing.T) {
		// Mocked model function returns an error.
		d.model.CreateFn = func(ctx context.Context, d *core.PurchaserProduct) error { return errors.New("could not create") }

		err := d.svc.CreatePurchaserProduct(context.Background(), d.PurchaserProduct1)
		require.Error(t, err)
		require.Contains(t, err.Error(), "could not create")
	})

}
