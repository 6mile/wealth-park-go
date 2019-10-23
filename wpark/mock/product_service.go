package mock

import (
	"context"

	"github.com/yashmurty/wealth-park/wpark/core"
)

// ProductService ...
type ProductService struct {
	CreateProductFn       func(ctx context.Context, d *core.Product) error
	CreateProductFnCalled int
}

var (
	_ core.ProductService = &ProductService{}
)

// CreateProduct ...
func (s *ProductService) CreateProduct(ctx context.Context, d *core.Product) error {
	s.CreateProductFnCalled++
	return s.CreateProductFn(ctx, d)
}
