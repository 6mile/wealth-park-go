package mock

import (
	"context"

	"github.com/yashmurty/wealth-park/wpark/core"
)

// PurchaserProductService ...
type PurchaserProductService struct {
	CreatePurchaserProductFn       func(ctx context.Context, d *core.PurchaserProduct) error
	CreatePurchaserProductFnCalled int
	ListPurchaserProductFn         func(ctx context.Context, purchaserID string, sArgs core.ListIncludeProductArgs) (all *core.ListPurchasesWithProductCustom, err error)
	ListPurchaserProductFnCalled   int
}

var (
	_ core.PurchaserProductService = &PurchaserProductService{}
)

// CreatePurchaserProduct ...
func (s *PurchaserProductService) CreatePurchaserProduct(ctx context.Context, d *core.PurchaserProduct) error {
	s.CreatePurchaserProductFnCalled++
	return s.CreatePurchaserProductFn(ctx, d)
}

// ListPurchaserProduct ...
func (s *PurchaserProductService) ListPurchaserProduct(ctx context.Context, purchaserID string, sArgs core.ListIncludeProductArgs) (all *core.ListPurchasesWithProductCustom, err error) {
	s.ListPurchaserProductFnCalled++
	return s.ListPurchaserProductFn(ctx, purchaserID, sArgs)
}
