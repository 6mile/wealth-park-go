package mysql

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yashmurty/wealth-park/wpark/core"
)

type ProductModelTestData struct {
	model        *ProductModel
	testProduct1 *core.Product
	testProduct2 *core.Product
}

func NewProductModelTestData() *ProductModelTestData {
	t := ProductModelTestData{}
	SetupDBHandle()

	t.testProduct1, _ = core.NewProduct(core.NewProductArgs{
		ID:   "PRODUCT-1",
		Name: "Test product 1 name",
	})

	t.testProduct2, _ = core.NewProduct(core.NewProductArgs{
		ID:   "PRODUCT-2",
		Name: "Test product 2 name",
	})

	t.model = NewProductModel()

	return &t
}

func TestProductCreateTable(t *testing.T) {
	d := NewProductModelTestData()

	t.Run("should succeed and create table", func(t *testing.T) {
		err := d.model.CreateTable(context.Background(), true)
		require.NoError(t, err)
	})
}

func TestProductCreate(t *testing.T) {
	d := NewProductModelTestData()

	t.Run("should succeed and create product", func(t *testing.T) {
		// Create runs successfully.
		err := d.model.Create(context.Background(), d.testProduct1)
		require.NoError(t, err)
	})

	t.Run("should fail due to duplicate primary key", func(t *testing.T) {
		// Create fails due to duplicate primary key.
		err := d.model.Create(context.Background(), d.testProduct1)
		require.Error(t, err)
	})

	t.Run("should succeed and create product", func(t *testing.T) {
		// Create runs successfully.
		err := d.model.Create(context.Background(), d.testProduct2)
		require.NoError(t, err)
	})

	t.Run("should fail due to duplicate name", func(t *testing.T) {
		// Create fails due to duplicate name.
		err := d.model.Create(context.Background(), d.testProduct2)
		require.Error(t, err)
	})

}
