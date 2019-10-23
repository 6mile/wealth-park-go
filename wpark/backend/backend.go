package backend

import (
	// Blank import the mysql driver.
	_ "github.com/go-sql-driver/mysql"

	"github.com/yashmurty/wealth-park/wpark/apiserver"
	con "github.com/yashmurty/wealth-park/wpark/controller"
	"github.com/yashmurty/wealth-park/wpark/core"
	"github.com/yashmurty/wealth-park/wpark/mysql"
	"github.com/yashmurty/wealth-park/wpark/pkg/util"
	"github.com/yashmurty/wealth-park/wpark/service"
)

// Backend instance.
type Backend struct {
	Server   *apiserver.Server
	Services Services
	Models   Models
}

// Services contains all initialized services used by the backend.
type Services struct {
	Product *service.ProductService
}

// Models contains all initialized models used by the backend.
type Models struct {
	Product core.ProductModel
}

// NewBackendWithMYSQLModels creates a backend using all mysql based
// model implementations.
func NewBackendWithMYSQLModels() *Backend {
	b := &Backend{}
	b.createMYSQLModels()
	b.wireServices()
	b.wireControllers()
	return b
}

func (b *Backend) createMYSQLModels() {
	// Setup models.
	mysql.SetupDBHandle()
	b.Models.Product = mysql.NewProductModel()
}

func (b *Backend) wireServices() {
	// Setup services.
	product := service.NewProductService()
	product.SetProductModel(b.Models.Product)

	b.Services.Product = product

	util.EnsureNoNilPointers(
		b.Services.Product,
	)
}

func (b *Backend) wireControllers() {
	// Setup controllers.
	con.ProductController.SetProductService(b.Services.Product)

	util.EnsureNoNilPointers(
		con.ProductController,
	)
}
