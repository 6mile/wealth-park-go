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
	testProduct3 *core.Product
	testProduct4 *core.Product
}

func NewProductModelTestData() *ProductModelTestData {
	t := ProductModelTestData{}
	SetupDBHandle()

	t.testProduct1, _ = core.NewProduct(core.NewProductArgs{
		ID:   "PRODUCT-1",
		Name: "Test product 1 name",
	})

	// Add another product which expires earlier than testProduct1.
	// testProduct2.EndDate < testProduct1.EndDate
	t.testProduct2, _ = core.NewProduct(core.NewProductArgs{
		ID:   "PRODUCT-2",
		Name: "Test product 2 name",
	})

	// Add another product which expires earlier than testProduct2.
	// testProduct3.EndDate < testProduct2.EndDate < testProduct1.EndDate
	t.testProduct3, _ = core.NewProduct(core.NewProductArgs{
		ID:   "PRODUCT-3",
		Name: "Test product 3 name",
	})

	// Add another product which expires earlier than testProduct3, but has a greater start date.
	// testProduct4.EndDate < testProduct3.EndDate < testProduct2.EndDate < testProduct1.EndDate
	t.testProduct4, _ = core.NewProduct(core.NewProductArgs{
		ID:   "PRODUCT-4",
		Name: "Test product 4 name",
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

	t.Run("should succeed and create product", func(t *testing.T) {
		// Create runs successfully.
		err := d.model.Create(context.Background(), d.testProduct3)
		require.NoError(t, err)
	})

	t.Run("should succeed and create product", func(t *testing.T) {
		// Create runs successfully.
		err := d.model.Create(context.Background(), d.testProduct4)
		require.NoError(t, err)
	})

}
