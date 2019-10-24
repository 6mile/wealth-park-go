package service

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yashmurty/wealth-park/wpark/core"
	"github.com/yashmurty/wealth-park/wpark/mock"
)

type PurchaserServiceTestData struct {
	svc        *PurchaserService
	model      *mock.PurchaserModel
	Purchaser1 *core.Purchaser
}

func NewPurchaserServiceTestData() *PurchaserServiceTestData {
	t := PurchaserServiceTestData{}

	t.Purchaser1, _ = core.NewPurchaser(core.NewPurchaserArgs{
		ID:   "PRODUCT-1",
		Name: "Test purchase 1 name",
	})

	t.svc = &PurchaserService{}

	t.model = &mock.PurchaserModel{}
	t.svc.SetPurchaserModel(t.model)

	return &t
}

func TestCreatePurchaser(t *testing.T) {
	d := NewPurchaserServiceTestData()

	t.Run("should succeed and create purchase", func(t *testing.T) {
		// Mocked model function runs successfully.
		d.model.CreateFn = func(ctx context.Context, d *core.Purchaser) error { return nil }

		err := d.svc.CreatePurchaser(context.Background(), d.Purchaser1)
		require.NoError(t, err)
		require.Equal(t, 1, d.model.CreateFnCalled)
	})

	t.Run("should fail since model function returns error", func(t *testing.T) {
		// Mocked model function returns an error.
		d.model.CreateFn = func(ctx context.Context, d *core.Purchaser) error { return errors.New("could not create") }

		err := d.svc.CreatePurchaser(context.Background(), d.Purchaser1)
		require.Error(t, err)
		require.Contains(t, err.Error(), "could not create")
	})

}
