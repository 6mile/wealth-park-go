package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yashmurty/wealth-park/wpark/core"
)

type productController struct {
	svc core.ProductService
}

// ProductController ...
var ProductController = &productController{}

func (s *productController) SetProductService(svc core.ProductService) { s.svc = svc }

// CreateProductRequestV1 is passed when creating a new product.
type CreateProductRequestV1 struct {
	Name string `json:"name" binding:"omitempty,min=1,max=2048"`
}

// CreateProductResponseV1 is returned when creating a new product.
type CreateProductResponseV1 struct {
	Product *core.Product `json:"product"`
}

func (s *productController) CreateProductV1(c *gin.Context) {
	r := CreateProductRequestV1{}
	err := c.BindJSON(&r)
	if err != nil {
		Fail(c, http.StatusBadRequest, err)
		return
	}

	// Create the product.
	b, err := core.NewProduct(core.NewProductArgs{
		Name: r.Name,
	})
	if err != nil {
		Fail(c, http.StatusBadRequest, err)
		return
	}

	err = s.svc.CreateProduct(c.Request.Context(), b)
	if err != nil {
		Fail(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, CreateProductResponseV1{
		Product: b,
	})
}

// ProductResponseV1 is returned when fetching a product.
type ProductResponseV1 struct {
	Product *core.Product `json:"product"`
}
