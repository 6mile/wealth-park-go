package core

import (
	"errors"

	"github.com/yashmurty/wealth-park/wpark/pkg/util"
)

/*
Product can be purchased.
*/
type Product struct {
	Resource
	Name string `json:"name"`
}

// NewProductArgs ...
type NewProductArgs struct {
	ID   string
	Name string
}

// NewProduct ...
func NewProduct(a NewProductArgs) (*Product, error) {
	if a.Name == "" {
		return nil, errors.New("product is missing name")
	}
	if a.ID == "" {
		a.ID = util.CreateID()
	}

	return &Product{
		Name: a.Name,
		Resource: Resource{
			ID:        a.ID,
			CreatedAt: util.MakeTimestamp(),
			UpdatedAt: util.MakeTimestamp(),
		},
	}, nil
}
