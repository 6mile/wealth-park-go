package mock

import (
	"context"

	"github.com/yashmurty/wealth-park/wpark/core"
)

// PurchaserModel ...
type PurchaserModel struct {
	CreateFn       func(ctx context.Context, d *core.Purchaser) error
	CreateFnCalled int
	BasicModel
}

var (
	_ core.PurchaserModel = &PurchaserModel{}
)

// Create ...
func (s *PurchaserModel) Create(ctx context.Context, d *core.Purchaser) error {
	s.CreateFnCalled++
	return s.CreateFn(ctx, d)
}
