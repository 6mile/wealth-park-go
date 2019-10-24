package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/yashmurty/wealth-park/wpark/core"
)

// PurchaserProductService ...
type PurchaserProductService struct {
	model core.PurchaserProductModel
	tag   string
}

var _ core.PurchaserProductService = &PurchaserProductService{}

// NewPurchaserProductService creates a new service instance.
func NewPurchaserProductService() *PurchaserProductService {
	return &PurchaserProductService{tag: "purchaser_product service"}
}

// SetPurchaserProductModel ...
func (s *PurchaserProductService) SetPurchaserProductModel(m core.PurchaserProductModel) { s.model = m }

// CreatePurchaserProduct ...
func (s *PurchaserProductService) CreatePurchaserProduct(ctx context.Context, b *core.PurchaserProduct) error {
	method := "create purchaser_product"

	// Create the purchaser_product.
	err := s.model.Create(ctx, b)
	if err != nil {
		return errors.Wrapf(err, serviceTag+": "+method+" failed in %s", s.tag)
	}

	return nil
}
