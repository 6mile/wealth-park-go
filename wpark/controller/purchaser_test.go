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

type PurchaserControllerTestData struct {
	c          *purchaserController
	svc        *mock.PurchaserService
	purchaser1 *core.Purchaser
}

func NewPurchaserControllerTestData() *PurchaserControllerTestData {
	apiserver.GetTestServer(SetupHTTPHandlers)

	t := PurchaserControllerTestData{}

	t.purchaser1, _ = core.NewPurchaser(core.NewPurchaserArgs{
		ID:   "PURCHASER-1",
		Name: "Test purchaser 1 name",
	})

	t.svc = &mock.PurchaserService{}

	t.c = PurchaserController
	t.c.SetPurchaserService(t.svc)

	return &t
}

func TestCreatePurchaserV1(t *testing.T) {
	d := NewPurchaserControllerTestData()

	in := CreatePurchaserRequestV1{
		Name: "Test purchaser 1 name",
	}
	out := CreatePurchaserResponseV1{}

	t.Run("should succeed and create purchaser", func(t *testing.T) {
		// Mocked service function runs successfully.
		d.svc.CreatePurchaserFn = func(ctx context.Context, b *core.Purchaser) error {
			require.Equal(t, d.purchaser1.Name, b.Name)
			return nil
		}

		w, _ := apiserver.CallAPI("POST", "/api/v1/purchaser", &in, &out)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, in.Name, out.Purchaser.Name)
		require.NotEmpty(t, out.Purchaser.ID)
	})

	t.Run("should fail since input request is invalid", func(t *testing.T) {
		bad := "Wait.. WUT?"
		w, _ := apiserver.CallAPI("POST", "/api/v1/purchaser", &bad, &out)
		require.Equal(t, 400, w.Code)
		require.Contains(t, w.Body.String(), "cannot unmarshal")
	})

	t.Run("should fail since service function returns error", func(t *testing.T) {
		// Mocked service function returns an error.
		d.svc.CreatePurchaserFn = func(ctx context.Context, wsa *core.Purchaser) error {
			return errors.New("Bad Request")
		}

		w, _ := apiserver.CallAPI("POST", "/api/v1/purchaser", &in, &out)
		require.Equal(t, 400, w.Code)
		require.Contains(t, w.Body.String(), "Bad Request")
	})

}
