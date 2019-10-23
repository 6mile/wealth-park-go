package e2e

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yashmurty/wealth-park/wpark/controller"
)

func TestCreateProduct(t *testing.T) {
	setupE2ETests()

	in := &controller.CreateProductRequestV1{
		Name: "E2E product name 1",
	}
	out := &controller.CreateProductResponseV1{}

	resp, _ := CallAPI("POST", "/api/v1/product", "", in, out)
	require.Equal(t, 200, resp.Code)

	require.Equal(t, in.Name, out.Product.Name)
	require.NotEmpty(t, out.Product.ID)
}
