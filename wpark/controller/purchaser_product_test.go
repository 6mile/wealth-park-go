package controller

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yashmurty/wealth-park/wpark/apiserver"
	"github.com/yashmurty/wealth-park/wpark/core"
	"github.com/yashmurty/wealth-park/wpark/mock"
)

type PurchaserProductControllerTestData struct {
	c                 *purchaserProductController
	svc               *mock.PurchaserProductService
	purchaserProduct1 *core.PurchaserProduct
}

func NewPurchaserProductControllerTestData() *PurchaserProductControllerTestData {
	apiserver.GetTestServer(SetupHTTPHandlers)

	t := PurchaserProductControllerTestData{}

	t.purchaserProduct1, _ = core.NewPurchaserProduct(core.NewPurchaserProductArgs{
		ID:                "PURCHASER-PRODUCT-1",
		PurchaserID:       "PURCHASER-1",
		ProductID:         "PRODUCT-1",
		PurchaseTimestamp: time.Now().Unix(),
	})

	t.svc = &mock.PurchaserProductService{}

	t.c = PurchaserProductController
	t.c.SetPurchaserProductService(t.svc)

	return &t
}

func TestCreatePurchaserProductV1(t *testing.T) {
	d := NewPurchaserProductControllerTestData()

	in := CreatePurchaserProductRequestV1{
		PurchaserID:       "PURCHASER-1",
		ProductID:         "PRODUCT-1",
		PurchaseTimestamp: time.Now().Unix(),
	}
	out := CreatePurchaserProductResponseV1{}

	t.Run("should succeed and create purchaser_product", func(t *testing.T) {
		// Mocked service function runs successfully.
		d.svc.CreatePurchaserProductFn = func(ctx context.Context, b *core.PurchaserProduct) error {
			require.Equal(t, d.purchaserProduct1.PurchaserID, b.PurchaserID)
			return nil
		}

		w, _ := apiserver.CallAPI("POST", "/api/v1/purchaser-product", &in, &out)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, in.PurchaserID, out.PurchaserProduct.PurchaserID)
		require.NotEmpty(t, out.PurchaserProduct.ID)
	})

	t.Run("should fail since input request is invalid", func(t *testing.T) {
		bad := "Wait.. WUT?"
		w, _ := apiserver.CallAPI("POST", "/api/v1/purchaser-product", &bad, &out)
		require.Equal(t, 400, w.Code)
		require.Contains(t, w.Body.String(), "cannot unmarshal")
	})

	t.Run("should fail since service function returns error", func(t *testing.T) {
		// Mocked service function returns an error.
		d.svc.CreatePurchaserProductFn = func(ctx context.Context, wsa *core.PurchaserProduct) error {
			return errors.New("Bad Request")
		}

		w, _ := apiserver.CallAPI("POST", "/api/v1/purchaser-product", &in, &out)
		require.Equal(t, 400, w.Code)
		require.Contains(t, w.Body.String(), "Bad Request")
	})

}

func TestListPurchaserProductV1(t *testing.T) {
	d := NewPurchaserProductControllerTestData()

	out := ListPurchaserProductResponseV1{}

	t.Run("should succeed and create purchaser_product", func(t *testing.T) {
		// Mocked service function runs successfully.
		query := url.Values{}
		query.Set("start_date", "2019-01-01")
		query.Set("end_date", "2019-12-31")
		querystring := query.Encode()

		d.svc.ListPurchaserProductFn = func(ctx context.Context, purchaserID string, sArgs core.ListIncludeProductArgs) (all []*core.PurchaserProduct, err error) {
			require.Equal(t, purchaserID, "WAT")
			return []*core.PurchaserProduct{d.purchaserProduct1}, nil
		}

		w, _ := apiserver.CallAPI("GET", "/api/v1/purchaser/WAT/product?"+querystring, nil, &out)
		require.Equal(t, http.StatusOK, w.Code)
		require.NotEmpty(t, out.PurchaserProducts)
	})

	t.Run("should fail since date format is wrong", func(t *testing.T) {
		query := url.Values{}
		query.Set("start_date", "1")
		query.Set("end_date", "2019-12-31")
		querystring := query.Encode()
		d.svc.ListPurchaserProductFn = func(ctx context.Context, purchaserID string, sArgs core.ListIncludeProductArgs) (all []*core.PurchaserProduct, err error) {
			require.Equal(t, purchaserID, "WAT")
			return []*core.PurchaserProduct{d.purchaserProduct1}, nil
		}

		w, _ := apiserver.CallAPI("GET", "/api/v1/purchaser/WAT/product?"+querystring, nil, &out)
		require.Equal(t, 400, w.Code)
		require.Contains(t, w.Body.String(), "Field validation")

		query = url.Values{}
		query.Set("start_date", "1234512345")
		query.Set("end_date", "2019-12-31")
		querystring = query.Encode()
		d.svc.ListPurchaserProductFn = func(ctx context.Context, purchaserID string, sArgs core.ListIncludeProductArgs) (all []*core.PurchaserProduct, err error) {
			require.Equal(t, purchaserID, "WAT")
			return []*core.PurchaserProduct{d.purchaserProduct1}, nil
		}

		w, _ = apiserver.CallAPI("GET", "/api/v1/purchaser/WAT/product?"+querystring, nil, &out)
		require.Equal(t, 400, w.Code)
		require.Contains(t, w.Body.String(), "cannot parse")

		query = url.Values{}
		query.Set("start_date", "2019-12-31")
		query.Set("end_date", "1234512345")
		querystring = query.Encode()
		d.svc.ListPurchaserProductFn = func(ctx context.Context, purchaserID string, sArgs core.ListIncludeProductArgs) (all []*core.PurchaserProduct, err error) {
			require.Equal(t, purchaserID, "WAT")
			return []*core.PurchaserProduct{d.purchaserProduct1}, nil
		}

		w, _ = apiserver.CallAPI("GET", "/api/v1/purchaser/WAT/product?"+querystring, nil, &out)
		require.Equal(t, 400, w.Code)
		require.Contains(t, w.Body.String(), "cannot parse")
	})

	t.Run("should fail since service function returns error", func(t *testing.T) {
		// Mocked service function returns an error.
		d.svc.ListPurchaserProductFn = func(ctx context.Context, purchaserID string, sArgs core.ListIncludeProductArgs) (all []*core.PurchaserProduct, err error) {
			require.Equal(t, purchaserID, "WAT")
			return []*core.PurchaserProduct{}, errors.New("Bad Request")
		}

		w, _ := apiserver.CallAPI("GET", "/api/v1/purchaser/WAT/product", nil, &out)
		require.Equal(t, 400, w.Code)
		require.Contains(t, w.Body.String(), "Bad Request")
	})
}
