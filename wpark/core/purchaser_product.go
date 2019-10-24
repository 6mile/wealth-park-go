package core

import (
	"context"
	"errors"

	"github.com/yashmurty/wealth-park/wpark/pkg/util"
)

/*
PurchaserProduct can purchase a product.
*/
type PurchaserProduct struct {
	Resource
	PurchaserID       string `json:"purchaser_id"`
	ProductID         string `json:"product_id"`
	PurchaseTimestamp int64  `json:"purchase_timestamp"`
}

// PurchaserProductModel describes data layer operations related to purchasers.
type PurchaserProductModel interface {
	Create(context.Context, *PurchaserProduct) error
	ListIncludeProduct(context.Context, string, ListIncludeProductArgs) ([]*PurchaserProduct, error)
	Model
}

// PurchaserProductService describes business logic operations related to purchasers.
type PurchaserProductService interface {
	CreatePurchaserProduct(context.Context, *PurchaserProduct) error
}

// NewPurchaserProductArgs ...
type NewPurchaserProductArgs struct {
	ID                string
	PurchaserID       string
	ProductID         string
	PurchaseTimestamp int64
}

// NewPurchaserProduct ...
func NewPurchaserProduct(a NewPurchaserProductArgs) (*PurchaserProduct, error) {
	if a.PurchaserID == "" {
		return nil, errors.New("purchaser_product is missing purchaser_id")
	}
	if a.ProductID == "" {
		return nil, errors.New("purchaser_product is missing product_id")
	}
	if a.PurchaseTimestamp == 0 {
		return nil, errors.New("purchaser_product is missing purchaser_timestamp")
	}
	if a.ID == "" {
		a.ID = util.CreateID()
	}

	return &PurchaserProduct{
		PurchaserID:       a.PurchaserID,
		ProductID:         a.ProductID,
		PurchaseTimestamp: a.PurchaseTimestamp,
		Resource: Resource{
			ID:        a.ID,
			CreatedAt: util.MakeTimestamp(),
			UpdatedAt: util.MakeTimestamp(),
		},
	}, nil
}

// ListIncludeProductArgs ...
type ListIncludeProductArgs struct {
	StartDate string `json:"start_date" binding:"omitempty,min=10,max=10"`
	EndDate   string `json:"end_date" binding:"omitempty,min=10,max=10"`
}
