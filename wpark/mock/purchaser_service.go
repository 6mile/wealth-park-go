package mock

import (
	"context"

	"github.com/yashmurty/wealth-park/wpark/core"
)

// PurchaserService ...
type PurchaserService struct {
	CreatePurchaserFn       func(ctx context.Context, d *core.Purchaser) error
	CreatePurchaserFnCalled int
}

var (
	_ core.PurchaserService = &PurchaserService{}
)

// CreatePurchaser ...
func (s *PurchaserService) CreatePurchaser(ctx context.Context, d *core.Purchaser) error {
	s.CreatePurchaserFnCalled++
	return s.CreatePurchaserFn(ctx, d)
}
