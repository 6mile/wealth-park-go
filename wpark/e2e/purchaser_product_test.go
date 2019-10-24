package e2e

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yashmurty/wealth-park/wpark/controller"
)

func TestCreatePurchaserProduct(t *testing.T) {
	setupE2ETests()

	in := &controller.CreatePurchaserProductRequestV1{
		PurchaserID:       "E2E-PURCHASER-1",
		ProductID:         "E2E-PRODUCT-1",
		PurchaseTimestamp: time.Now().Unix(),
	}
	out := &controller.CreatePurchaserProductResponseV1{}

	resp, _ := CallAPI("POST", "/api/v1/purchaser-product", "", in, out)
	require.Equal(t, 200, resp.Code)

	require.Equal(t, in.PurchaserID, out.PurchaserProduct.PurchaserID)
	require.NotEmpty(t, out.PurchaserProduct.ID)
}
