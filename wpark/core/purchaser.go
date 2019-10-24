package core

import (
	"context"
	"errors"

	"github.com/yashmurty/wealth-park/wpark/pkg/util"
)

/*
Purchaser can purchase a product.
*/
type Purchaser struct {
	Resource
	Name string `json:"name"`
}

// PurchaserModel describes data layer operations related to purchasers.
type PurchaserModel interface {
	Create(context.Context, *Purchaser) error
	Model
}

// PurchaserService describes business logic operations related to purchasers.
type PurchaserService interface {
	CreatePurchaser(context.Context, *Purchaser) error
}

// NewPurchaserArgs ...
type NewPurchaserArgs struct {
	ID   string
	Name string
}

// NewPurchaser ...
func NewPurchaser(a NewPurchaserArgs) (*Purchaser, error) {
	if a.Name == "" {
		return nil, errors.New("purchaser is missing name")
	}
	if a.ID == "" {
		a.ID = util.CreateID()
	}

	return &Purchaser{
		Name: a.Name,
		Resource: Resource{
			ID:        a.ID,
			CreatedAt: util.MakeTimestamp(),
			UpdatedAt: util.MakeTimestamp(),
		},
	}, nil
}
