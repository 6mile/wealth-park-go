package controller

import (
	"net/http"
	"time"

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

// ListPurchaserProductRequestV1 is passed when creating a new purchaser_product.
type ListPurchaserProductRequestV1 struct {
	StartDate string `json:"start_date" form:"start_date" binding:"omitempty,min=10,max=10"`
	EndDate   string `json:"end_date" form:"end_date" binding:"omitempty,min=10,max=10"`
}

// ListPurchaserProductResponseV1 is returned when creating a new purchaser_product.
type ListPurchaserProductResponseV1 struct {
	PurchaserProducts []*core.PurchaserProduct `json:"purchaser_products"`
}

func (s *purchaserProductController) ListPurchaserProductV1(c *gin.Context) {
	purchaserID := c.Param("purchaser_id")

	r := ListPurchaserProductRequestV1{}
	err := c.BindQuery(&r)
	if err != nil {
		Fail(c, http.StatusBadRequest, err)
		return
	}

	// Convert date to timestamp.
	dateLayout := "2006-01-02"
	var startDateTimestamp int64
	if r.StartDate != "" {
		t, err := time.Parse(dateLayout, r.StartDate)
		if err != nil {
			Fail(c, http.StatusBadRequest, err)
			return
		}
		startDateTimestamp = t.Unix()
	}
	var endDateTimestamp int64
	if r.EndDate != "" {
		t, err := time.Parse(dateLayout, r.EndDate)
		if err != nil {
			Fail(c, http.StatusBadRequest, err)
			return
		}
		endDateTimestamp = t.Unix()
	}
	sArgs := core.ListIncludeProductArgs{
		StartDateTimestamp: startDateTimestamp,
		EndDateTimestamp:   endDateTimestamp,
	}
	all, err := s.svc.ListPurchaserProduct(c.Request.Context(), purchaserID, sArgs)
	if err != nil {
		Fail(c, http.StatusBadRequest, err)
		return
	}

	resp := ListPurchaserProductResponseV1{PurchaserProducts: all}
	if resp.PurchaserProducts == nil {
		resp.PurchaserProducts = make([]*core.PurchaserProduct, 0)
	}
	c.JSON(http.StatusOK, resp)

}
