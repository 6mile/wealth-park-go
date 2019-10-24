package e2e

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yashmurty/wealth-park/wpark/controller"
)

func TestCreatePurchaserProduct(t *testing.T) {
	setupE2ETests()

	testPurchaser := createTestPurchaserData("E2E - TestCreatePurchaserProduct purchaser")
	testProduct := createTestProductData("E2E - TestCreatePurchaserProduct product")

	in := &controller.CreatePurchaserProductRequestV1{
		PurchaserID:       testPurchaser.ID,
		ProductID:         testProduct.ID,
		PurchaseTimestamp: time.Now().Unix(),
	}
	out := &controller.CreatePurchaserProductResponseV1{}

	resp, _ := CallAPI("POST", "/api/v1/purchaser-product", "", in, out)
	require.Equal(t, 200, resp.Code)

	require.Equal(t, in.PurchaserID, out.PurchaserProduct.PurchaserID)
	require.NotEmpty(t, out.PurchaserProduct.ID)
}

func TestListPurchaserProduct(t *testing.T) {
	setupE2ETests()

	// Create dependant data.
	testPurchaser1 := createTestPurchaserData("E2E - TestListPurchaserProduct purchaser 1")
	testProduct1 := createTestProductData("E2E - TestListPurchaserProduct product 1")
	testProduct2 := createTestProductData("E2E - TestListPurchaserProduct product 2")

	in := &controller.CreatePurchaserProductRequestV1{
		PurchaserID:       testPurchaser1.ID,
		ProductID:         testProduct1.ID,
		PurchaseTimestamp: time.Now().Unix(),
	}
	out := &controller.CreatePurchaserProductResponseV1{}

	resp, _ := CallAPI("POST", "/api/v1/purchaser-product", "", in, out)
	require.Equal(t, 200, resp.Code)

	in = &controller.CreatePurchaserProductRequestV1{
		PurchaserID:       testPurchaser1.ID,
		ProductID:         testProduct2.ID,
		PurchaseTimestamp: time.Now().AddDate(0, 0, -20).Unix(),
	}
	resp, _ = CallAPI("POST", "/api/v1/purchaser-product", "", in, out)
	require.Equal(t, 200, resp.Code)

	out2 := &controller.ListPurchaserProductResponseV1{}

	// Call the api.
	resp, _ = CallAPI("GET", "/api/v1/purchaser/"+testPurchaser1.ID+"/product", "", nil, out2)
	require.Equal(t, 200, resp.Code)

	require.Equal(t, http.StatusOK, resp.Code)
	require.Equal(t, 2, len(out2.PurchaserProducts))

	query := url.Values{}
	query.Set("start_date", time.Now().AddDate(0, 0, -10).Format("2006-01-02"))
	query.Set("end_date", time.Now().AddDate(0, 0, 10).Format("2006-01-02"))
	querystring := query.Encode()

	// Call the api.
	resp, _ = CallAPI("GET", "/api/v1/purchaser/"+testPurchaser1.ID+"/product?"+querystring, "", nil, out2)
	require.Equal(t, 200, resp.Code)
	require.Equal(t, 1, len(out2.PurchaserProducts))
}
