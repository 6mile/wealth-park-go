package mysql

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yashmurty/wealth-park/wpark/core"
)

type PurchaserProductModelTestData struct {
	model                 *PurchaserProductModel
	testPurchaser1        *core.Purchaser
	testProduct1          *core.Product
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

	t.Run("should succeed and create purchaser", func(t *testing.T) {
		// Create runs successfully.
		fmt.Println("d.testPurchaserProduct1 : ", d.testPurchaserProduct1)
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

func setupPurchaserProductDependentData(t *PurchaserProductModelTestData) {
	t.testPurchaser1, _ = core.NewPurchaser(core.NewPurchaserArgs{
		Name: "Setup purchaser 1 name",
	})
	err := NewPurchaserModel().CreateTable(context.Background(), true)
	if err != nil {
		panic(err)
	}

	err = NewPurchaserModel().Create(context.Background(), t.testPurchaser1)
	if err != nil {
		panic(err)
	}

	t.testProduct1, _ = core.NewProduct(core.NewProductArgs{
		Name: "Setup product 1 name",
	})
	err = NewProductModel().CreateTable(context.Background(), true)
	if err != nil {
		panic(err)
	}
	err = NewProductModel().Create(context.Background(), t.testProduct1)
	if err != nil {
		panic(err)
	}
}
