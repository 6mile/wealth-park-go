package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/yashmurty/wealth-park/wpark/core"
)

// ProductService ...
type ProductService struct {
	model core.ProductModel
	tag   string
}

var _ core.ProductService = &ProductService{}

// NewProductService creates a new service instance.
func NewProductService() *ProductService {
	return &ProductService{tag: "product service"}
}

// SetProductModel ...
func (s *ProductService) SetProductModel(m core.ProductModel) { s.model = m }

// CreateProduct ...
func (s *ProductService) CreateProduct(ctx context.Context, b *core.Product) error {
	method := "create product"

	// Create the product.
	err := s.model.Create(ctx, b)
	if err != nil {
		return errors.Wrapf(err, serviceTag+": "+method+" failed in %s", s.tag)
	}

	return nil
}
