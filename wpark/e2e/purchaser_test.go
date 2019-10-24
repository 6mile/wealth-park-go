package e2e

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yashmurty/wealth-park/wpark/controller"
)

func TestCreatePurchaser(t *testing.T) {
	setupE2ETests()

	in := &controller.CreatePurchaserRequestV1{
		Name: "E2E purchaser name 1",
	}
	out := &controller.CreatePurchaserResponseV1{}

	resp, _ := CallAPI("POST", "/api/v1/purchaser", "", in, out)
	require.Equal(t, 200, resp.Code)

	require.Equal(t, in.Name, out.Purchaser.Name)
	require.NotEmpty(t, out.Purchaser.ID)
}
