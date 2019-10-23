package controller

import (
	"github.com/yashmurty/wealth-park/wpark/core"
)

type exampleData struct {
	Product *core.Product
}

// ExampleData contains example data used in our auto-generated
// API documentation.
var ExampleData = &exampleData{}

func init() {
	d := ExampleData
	var err error
	noError := func() {
		if err != nil {
			panic(err)
		}
	}

	d.Product, err = core.NewProduct(core.NewProductArgs{
		ID:   "PRODUCT-1",
		Name: "Test product name",
	})
	noError()
}
