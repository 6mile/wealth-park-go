package apiserver

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/yashmurty/wealth-park/wpark/pkg/logger"
	"go.uber.org/zap"
)

var (
	log = logger.Get("apiserver")
)

// Server represents an HTTPS (TLS) web server.
// It has a Gin engine, public routes and private routes
// that require a user session.
type Server struct {
	Engine       *gin.Engine
	PublicRouter *gin.RouterGroup
	Addr         string
	ServerID     string
}

// NewServerArgs is passed to NewServer.
type NewServerArgs struct {
	ServerID string
	Addr     string
}

func newStandardGinEngine() *gin.Engine {
	// Setup default router.
	e := gin.Default()

	// CORS all the things.
	e.Use(CORSMiddleware())

	// Compress all the things.
	e.Use(gzip.Gzip(gzip.DefaultCompression))

	// Good security defaults.
	e.Use(helmet.Default())

	e.Use(MetricsMiddleware())

	// This will print the error and stacktrace to stderr.
	e.Use(gin.Recovery())

	return e
}

// NewServer creates a new API server with suitable middleware
// and some basic routes for testing.
func NewServer(a NewServerArgs) *Server {
	// Force the use of validator.v9.
	binding.Validator = new(defaultValidator)

	// Setup default router.
	e := newStandardGinEngine()

	// Create route groups.
	pub := e.Group("")

	// Root route.
	pub.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, NewAPIStatus(200, "we’re good."))
	})

	pub.GET("/fail", func(c *gin.Context) {
		APIError(c, http.StatusInternalServerError, errors.New("all hell broke loose"))
	})

	return &Server{
		ServerID:     a.ServerID,
		Engine:       e,
		PublicRouter: pub,
		Addr:         a.Addr,
	}
}

// Start the HTTPS server.
func (s *Server) Start() {
	srv := &http.Server{
		Addr:    s.Addr,
		Handler: s.Engine,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("run failed", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	timeout := 5
	log.Info("shutting down server", zap.Int("timeout", timeout))

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("graceful server shutdown failed", zap.Error(err))
	}

	log.Info("server exiting - so long, fellas")
}

// APIError returns a JSON error with useful metadata.
func APIError(c *gin.Context, code int, err error) {
	if code >= http.StatusInternalServerError {
		log.Error(
			"api error",
			zap.String("path", c.Request.RequestURI),
			zap.Error(err),
		)
	}
	c.AbortWithStatusJSON(code, NewAPIStatus(code, err.Error()))
}

// APIStatus represents a standard JSON response when calling
// endpoints.
type APIStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewAPIStatus creates a new APIStatus.
func NewAPIStatus(code int, message string) APIStatus {
	return APIStatus{
		Code:    code,
		Message: message,
	}
}
