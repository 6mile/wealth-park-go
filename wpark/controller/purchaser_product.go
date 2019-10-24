package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yashmurty/wealth-park/wpark/core"
)

type purchaserProductController struct {
	svc core.PurchaserProductService
}

// PurchaserProductController ...
var PurchaserProductController = &purchaserProductController{}

func (s *purchaserProductController) SetPurchaserProductService(svc core.PurchaserProductService) {
	s.svc = svc
}

// CreatePurchaserProductRequestV1 is passed when creating a new purchaser_product.
type CreatePurchaserProductRequestV1 struct {
	PurchaserID       string `json:"purchaser_id" binding:"required,min=1,max=2048"`
	ProductID         string `json:"product_id" binding:"required,min=1,max=2048"`
	PurchaseTimestamp int64  `json:"purchase_timestamp" binding:"required,gt=0"`
}

// CreatePurchaserProductResponseV1 is returned when creating a new purchaser_product.
type CreatePurchaserProductResponseV1 struct {
	PurchaserProduct *core.PurchaserProduct `json:"purchaser_product"`
}

func (s *purchaserProductController) CreatePurchaserProductV1(c *gin.Context) {
	r := CreatePurchaserProductRequestV1{}
	err := c.BindJSON(&r)
	if err != nil {
		Fail(c, http.StatusBadRequest, err)
		return
	}

	// Create the purchaser_product.
	b, err := core.NewPurchaserProduct(core.NewPurchaserProductArgs{
		PurchaserID:       r.PurchaserID,
		ProductID:         r.ProductID,
		PurchaseTimestamp: r.PurchaseTimestamp,
	})
	if err != nil {
		Fail(c, http.StatusBadRequest, err)
		return
	}

	err = s.svc.CreatePurchaserProduct(c.Request.Context(), b)
	if err != nil {
		Fail(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, CreatePurchaserProductResponseV1{
		PurchaserProduct: b,
	})
}
