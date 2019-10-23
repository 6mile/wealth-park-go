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

// SetupEndpoint adds an endpoint to the given API server.
func SetupEndpoint(s *apiserver.Server, e Endpoint) {
	s.PublicRouter.Handle(e.Method, e.RelativePath, e.Handler)
}
