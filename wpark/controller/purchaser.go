package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yashmurty/wealth-park/wpark/core"
)

type purchaserController struct {
	svc core.PurchaserService
}

// PurchaserController ...
var PurchaserController = &purchaserController{}

func (s *purchaserController) SetPurchaserService(svc core.PurchaserService) { s.svc = svc }

// CreatePurchaserRequestV1 is passed when creating a new purchaser.
type CreatePurchaserRequestV1 struct {
	Name string `json:"name" binding:"omitempty,min=1,max=2048"`
}

// CreatePurchaserResponseV1 is returned when creating a new purchaser.
type CreatePurchaserResponseV1 struct {
	Purchaser *core.Purchaser `json:"purchaser"`
}

func (s *purchaserController) CreatePurchaserV1(c *gin.Context) {
	r := CreatePurchaserRequestV1{}
	err := c.BindJSON(&r)
	if err != nil {
		Fail(c, http.StatusBadRequest, err)
		return
	}

	// Create the purchaser.
	b, err := core.NewPurchaser(core.NewPurchaserArgs{
		Name: r.Name,
	})
	if err != nil {
		Fail(c, http.StatusBadRequest, err)
		return
	}

	err = s.svc.CreatePurchaser(c.Request.Context(), b)
	if err != nil {
		Fail(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, CreatePurchaserResponseV1{
		Purchaser: b,
	})
}

// PurchaserResponseV1 is returned when fetching a purchaser.
type PurchaserResponseV1 struct {
	Purchaser *core.Purchaser `json:"purchaser"`
}
