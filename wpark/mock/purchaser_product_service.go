package mock

import (
	"context"

	"github.com/yashmurty/wealth-park/wpark/core"
)

// PurchaserProductService ...
type PurchaserProductService struct {
	CreatePurchaserProductFn       func(ctx context.Context, d *core.PurchaserProduct) error
	CreatePurchaserProductFnCalled int
}

var (
	_ core.PurchaserProductService = &PurchaserProductService{}
)

// CreatePurchaserProduct ...
func (s *PurchaserProductService) CreatePurchaserProduct(ctx context.Context, d *core.PurchaserProduct) error {
	s.CreatePurchaserProductFnCalled++
	return s.CreatePurchaserProductFn(ctx, d)
}
