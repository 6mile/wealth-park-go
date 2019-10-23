package controller

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yashmurty/wealth-park/wpark/apiserver"
	"github.com/yashmurty/wealth-park/wpark/core"
	"github.com/yashmurty/wealth-park/wpark/mock"
)

type ProductControllerTestData struct {
	c        *productController
	svc      *mock.ProductService
	Product1 *core.Product
}

func NewProductControllerTestData() *ProductControllerTestData {
	_ = apiserver.GetTestServer(func(s *apiserver.Server) {
		for _, e := range ProductController.Endpoints() {
			SetupEndpoint(s, e)
		}
	})

	t := ProductControllerTestData{}

	t.Product1, _ = core.NewProduct(core.NewProductArgs{
		ID:   "PRODUCT-1",
		Name: "Test product 1 name",
	})

	t.svc = &mock.ProductService{}

	t.c = ProductController
	t.c.SetProductService(t.svc)

	return &t
}

func TestCreateProductV1(t *testing.T) {
	d := NewProductControllerTestData()

	in := CreateProductRequestV1{
		Name: "Test product 1 name",
	}
	out := CreateProductResponseV1{}

	t.Run("should succeed and create product", func(t *testing.T) {
		// Mocked service function runs successfully.
		d.svc.CreateProductFn = func(ctx context.Context, b *core.Product) error {
			require.Equal(t, d.Product1.Name, b.Name)
			return nil
		}

		w, _ := apiserver.CallAPI("POST", "/api/v1/product", &in, &out)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, in.Name, out.Product.Name)
		require.NotEmpty(t, out.Product.ID)
	})

	t.Run("should fail since input request is invalid", func(t *testing.T) {
		bad := "Wait.. WUT?"
		w, _ := apiserver.CallAPI("POST", "/api/v1/product", &bad, &out)
		require.Equal(t, 400, w.Code)
		require.Contains(t, w.Body.String(), "cannot unmarshal")
	})

	t.Run("should fail since service function returns error", func(t *testing.T) {
		// Mocked service function returns an error.
		d.svc.CreateProductFn = func(ctx context.Context, wsa *core.Product) error {
			return errors.New("Bad Request")
		}

		w, _ := apiserver.CallAPI("POST", "/api/v1/product", &in, &out)
		require.Equal(t, 400, w.Code)
		require.Contains(t, w.Body.String(), "Bad Request")
	})

}
