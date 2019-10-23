package service

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yashmurty/wealth-park/wpark/core"
	"github.com/yashmurty/wealth-park/wpark/mock"
)

type ProductServiceTestData struct {
	svc      *ProductService
	model    *mock.ProductModel
	Product1 *core.Product
}

func NewProductServiceTestData() *ProductServiceTestData {
	t := ProductServiceTestData{}

	t.Product1, _ = core.NewProduct(core.NewProductArgs{
		ID:   "PRODUCT-1",
		Name: "Test product 1 name",
	})

	t.svc = &ProductService{}

	t.model = &mock.ProductModel{}
	t.svc.SetProductModel(t.model)

	return &t
}

func TestCreateProduct(t *testing.T) {
	d := NewProductServiceTestData()

	t.Run("should succeed and create product", func(t *testing.T) {
		// Mocked model function runs successfully.
		d.model.CreateFn = func(ctx context.Context, d *core.Product) error { return nil }

		err := d.svc.CreateProduct(context.Background(), d.Product1)
		require.NoError(t, err)
		require.Equal(t, 1, d.model.CreateFnCalled)
	})

	t.Run("should fail since model function returns error", func(t *testing.T) {
		// Mocked model function returns an error.
		d.model.CreateFn = func(ctx context.Context, d *core.Product) error { return errors.New("could not create") }

		err := d.svc.CreateProduct(context.Background(), d.Product1)
		require.Error(t, err)
		require.Contains(t, err.Error(), "could not create")
	})

}
