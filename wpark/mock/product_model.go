package mock

import (
	"context"

	"github.com/yashmurty/wealth-park/wpark/core"
)

// ProductModel ...
type ProductModel struct {
	CreateFn       func(ctx context.Context, d *core.Product) error
	CreateFnCalled int
	BasicModel
}

var (
	_ core.ProductModel = &ProductModel{}
)

// Create ...
func (s *ProductModel) Create(ctx context.Context, d *core.Product) error {
	s.CreateFnCalled++
	return s.CreateFn(ctx, d)
}
