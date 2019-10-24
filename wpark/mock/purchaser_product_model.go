package mock

import (
	"context"

	"github.com/yashmurty/wealth-park/wpark/core"
)

// PurchaserProductModel ...
type PurchaserProductModel struct {
	CreateFn       func(ctx context.Context, d *core.PurchaserProduct) error
	CreateFnCalled int
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
