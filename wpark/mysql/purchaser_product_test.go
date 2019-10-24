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
	testPurchaser1        *core.Purchaser
	testPurchaser2        *core.Purchaser
	testPurchaser3        *core.Purchaser
	testProduct1          *core.Product
	testProduct2          *core.Product
	testProduct3          *core.Product
	testPurchaserProduct1 *core.PurchaserProduct
}

func NewPurchaserProductModelTestData() *PurchaserProductModelTestData {
	t := PurchaserProductModelTestData{}
	SetupDBHandle()

	setupPurchaserProductDependentData(&t)

	t.testPurchaserProduct1, _ = core.NewPurchaserProduct(core.NewPurchaserProductArgs{
		ID:                "PURCHASER-PRODUCT-1",
		PurchaserID:       t.testPurchaser1.ID,
		ProductID:         t.testProduct1.ID,
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

	t.Run("should succeed and create purchaser_product", func(t *testing.T) {
		// Create runs successfully.
		err := d.model.Create(context.Background(), d.testPurchaserProduct1)
		require.NoError(t, err)
	})

	t.Run("should fail due to duplicate primary key", func(t *testing.T) {
		// Create fails due to duplicate primary key.
		err := d.model.Create(context.Background(), d.testPurchaserProduct1)
		require.Error(t, err)
	})

	t.Run("should fail due to duplicate name", func(t *testing.T) {
		// Create fails due to duplicate name.
		err := d.model.Create(context.Background(), d.testPurchaserProduct1)
		require.Error(t, err)
	})

}

func TestPurchaserProductListIncludeProduct(t *testing.T) {
	d := NewPurchaserProductModelTestData()

	t.Run("should succeed and create purchaser_product", func(t *testing.T) {
		testPurchaserProduct2, _ := core.NewPurchaserProduct(core.NewPurchaserProductArgs{
			ID:                "PURCHASER-PRODUCT-2",
			PurchaserID:       d.testPurchaser1.ID,
			ProductID:         d.testProduct2.ID,
			PurchaseTimestamp: time.Now().Unix(),
		})
		// Create runs successfully.
		err := d.model.Create(context.Background(), testPurchaserProduct2)
		require.NoError(t, err)

		testPurchaserProduct3, _ := core.NewPurchaserProduct(core.NewPurchaserProductArgs{
			ID:                "PURCHASER-PRODUCT-3",
			PurchaserID:       d.testPurchaser3.ID,
			ProductID:         d.testProduct3.ID,
			PurchaseTimestamp: time.Now().Unix(),
		})
		// Create runs successfully.
		err = d.model.Create(context.Background(), testPurchaserProduct3)
		require.NoError(t, err)
	})

	t.Run("should succeed and list purchaser", func(t *testing.T) {
		// Create runs successfully.
		all, err := d.model.ListIncludeProduct(context.Background())
		require.Equal(t, 3, len(all))
		require.NoError(t, err)
	})
}

func setupPurchaserProductDependentData(t *PurchaserProductModelTestData) {
	err := NewPurchaserModel().CreateTable(context.Background(), true)
	if err != nil {
		panic(err)
	}

	t.testPurchaser1, _ = core.NewPurchaser(core.NewPurchaserArgs{
		Name: "Setup purchaser 1 name",
	})
	err = NewPurchaserModel().Create(context.Background(), t.testPurchaser1)
	if err != nil {
		panic(err)
	}
	t.testPurchaser2, _ = core.NewPurchaser(core.NewPurchaserArgs{
		Name: "Setup purchaser 2 name",
	})
	err = NewPurchaserModel().Create(context.Background(), t.testPurchaser2)
	if err != nil {
		panic(err)
	}
	t.testPurchaser3, _ = core.NewPurchaser(core.NewPurchaserArgs{
		Name: "Setup purchaser 3 name",
	})
	err = NewPurchaserModel().Create(context.Background(), t.testPurchaser3)
	if err != nil {
		panic(err)
	}

	err = NewProductModel().CreateTable(context.Background(), true)
	if err != nil {
		panic(err)
	}

	t.testProduct1, _ = core.NewProduct(core.NewProductArgs{
		Name: "Setup product 1 name",
	})
	err = NewProductModel().Create(context.Background(), t.testProduct1)
	if err != nil {
		panic(err)
	}
	t.testProduct2, _ = core.NewProduct(core.NewProductArgs{
		Name: "Setup product 2 name",
	})
	err = NewProductModel().Create(context.Background(), t.testProduct2)
	if err != nil {
		panic(err)
	}
	t.testProduct3, _ = core.NewProduct(core.NewProductArgs{
		Name: "Setup product 3 name",
	})
	err = NewProductModel().Create(context.Background(), t.testProduct3)
	if err != nil {
		panic(err)
	}
}
