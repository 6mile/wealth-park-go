package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yashmurty/wealth-park/wpark/core"
	"github.com/yashmurty/wealth-park/wpark/mock"
)

type PurchaserProductServiceTestData struct {
	svc                   *PurchaserProductService
	model                 *mock.PurchaserProductModel
	purchaserProduct1     *core.PurchaserProduct
	listPurchaserProduct1 *core.ListPurchaserProduct
}

func NewPurchaserProductServiceTestData() *PurchaserProductServiceTestData {
	t := PurchaserProductServiceTestData{}

	t.purchaserProduct1, _ = core.NewPurchaserProduct(core.NewPurchaserProductArgs{
		ID:                "PRODUCT-1",
		PurchaserID:       "PURCHASER-1",
		ProductID:         "PRODUCT-1",
		PurchaseTimestamp: time.Now().Unix(),
	})

	t.listPurchaserProduct1 = &core.ListPurchaserProduct{
		PurchaserProduct: *t.purchaserProduct1,
		ProductName:      "Test product name 1",
		DateOnly:         "2019-01-02",
	}

	t.svc = &PurchaserProductService{}

	t.model = &mock.PurchaserProductModel{}
	t.svc.SetPurchaserProductModel(t.model)

	return &t
}

func TestCreatePurchaserProduct(t *testing.T) {
	d := NewPurchaserProductServiceTestData()

	t.Run("should succeed and create purchase_product", func(t *testing.T) {
		// Mocked model function runs successfully.
		d.model.CreateFn = func(ctx context.Context, d *core.PurchaserProduct) error { return nil }

		err := d.svc.CreatePurchaserProduct(context.Background(), d.purchaserProduct1)
		require.NoError(t, err)
		require.Equal(t, 1, d.model.CreateFnCalled)
	})

	t.Run("should fail since model function returns error", func(t *testing.T) {
		// Mocked model function returns an error.
		d.model.CreateFn = func(ctx context.Context, d *core.PurchaserProduct) error { return errors.New("could not create") }

		err := d.svc.CreatePurchaserProduct(context.Background(), d.purchaserProduct1)
		require.Error(t, err)
		require.Contains(t, err.Error(), "could not create")
	})

}

func TestListPurchaserProduct(t *testing.T) {
	d := NewPurchaserProductServiceTestData()

	t.Run("should succeed and list purchase_product", func(t *testing.T) {
		// Mocked model function runs successfully.
		d.model.ListIncludeProductFn = func(ctx context.Context, purchaserID string, sArgs core.ListIncludeProductArgs) ([]*core.ListPurchaserProduct, error) {
			return []*core.ListPurchaserProduct{d.listPurchaserProduct1}, nil
		}

		fmt.Println("d.listPurchaserProduct1 : ", d.listPurchaserProduct1)

		list, err := d.svc.ListPurchaserProduct(
			context.Background(),
			d.purchaserProduct1.PurchaserID,
			core.ListIncludeProductArgs{},
		)
		require.NoError(t, err)
		require.Equal(t, 1, d.model.ListIncludeProductFnCalled)
		require.NotEmpty(t, list)
		require.NotEmpty(t, list.Purchases["2019-01-02"])
	})

	t.Run("should fail since model function returns error", func(t *testing.T) {
		// Mocked model function returns an error.
		d.model.ListIncludeProductFn = func(ctx context.Context, purchaserID string, sArgs core.ListIncludeProductArgs) ([]*core.ListPurchaserProduct, error) {
			return nil, errors.New("could not list")
		}

		_, err := d.svc.ListPurchaserProduct(
			context.Background(),
			d.purchaserProduct1.PurchaserID,
			core.ListIncludeProductArgs{},
		)

		require.Error(t, err)
		require.Contains(t, err.Error(), "could not list")
	})
}
