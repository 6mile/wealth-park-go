package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/yashmurty/wealth-park/wpark/core"
)

// PurchaserService ...
type PurchaserService struct {
	model core.PurchaserModel
	tag   string
}

var _ core.PurchaserService = &PurchaserService{}

// NewPurchaserService creates a new service instance.
func NewPurchaserService() *PurchaserService {
	return &PurchaserService{tag: "purchaser service"}
}

// SetPurchaserModel ...
func (s *PurchaserService) SetPurchaserModel(m core.PurchaserModel) { s.model = m }

// CreatePurchaser ...
func (s *PurchaserService) CreatePurchaser(ctx context.Context, b *core.Purchaser) error {
	method := "create purchaser"

	// Create the purchaser.
	err := s.model.Create(ctx, b)
	if err != nil {
		return errors.Wrapf(err, serviceTag+": "+method+" failed in %s", s.tag)
	}

	return nil
}
