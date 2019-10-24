package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yashmurty/wealth-park/wpark/apiserver"
	"github.com/yashmurty/wealth-park/wpark/pkg/logger"
)

var (
	log = logger.Get("controller")
)

// Endpoint implements a handler and is connected to route.
type Endpoint struct {
	Method       string             `json:"method"`
	RelativePath string             `json:"relative_path"`
	Description  string             `json:"description"`
	Request      interface{}        `json:"request"`
	Response     interface{}        `json:"response"`
	Handler      func(*gin.Context) `json:"-"`
	IsPublic     bool               `json:"is_public"`
}

// Empty is an empty endpoint response.
type Empty struct{}

// Fail returns an error response to an API call.
func Fail(c *gin.Context, code int, err error) {
	apiserver.APIError(c, code, err)
}

// SetupHTTPHandlers sets up http handlers on the given api server.
func SetupHTTPHandlers(s *apiserver.Server) {
	s.PublicRouter.Handle("POST", "/api/v1/product", ProductController.CreateProductV1)
	s.PublicRouter.Handle("POST", "/api/v1/purchaser", PurchaserController.CreatePurchaserV1)
	s.PublicRouter.Handle("POST", "/api/v1/purchaser-product", PurchaserProductController.CreatePurchaserProductV1)
	s.PublicRouter.Handle("GET", "/api/v1/purchaser/:purchaser_id/product", PurchaserProductController.ListPurchaserProductV1)
}
