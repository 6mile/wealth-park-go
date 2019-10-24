package mock

import (
	"context"

	"github.com/yashmurty/wealth-park/wpark/core"
)

// PurchaserProductModel ...
type PurchaserProductModel struct {
	CreateFn                   func(ctx context.Context, d *core.PurchaserProduct) error
	CreateFnCalled             int
	ListIncludeProductFn       func(ctx context.Context, purchaserID string, sArgs core.ListIncludeProductArgs) ([]*core.PurchaserProduct, error)
	ListIncludeProductFnCalled int
	BasicModel
}

var (
	_ core.PurchaserProductModel = &PurchaserProductModel{}
)

// Create ...
func (s *PurchaserProductModel) Create(ctx context.Context, d *core.PurchaserProduct) error {
	s.CreateFnCalled++
	return s.CreateFn(ctx, d)
}

// ListIncludeProduct ...
func (s *PurchaserProductModel) ListIncludeProduct(ctx context.Context, purchaserID string, sArgs core.ListIncludeProductArgs) ([]*core.PurchaserProduct, error) {
	s.ListIncludeProductFnCalled++
	return s.ListIncludeProductFn(ctx, purchaserID, sArgs)
}
