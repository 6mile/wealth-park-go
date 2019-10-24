package backend

import (
	// Blank import the mysql driver.
	"context"

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
	Product   *service.ProductService
	Purchaser *service.PurchaserService
}

// Models contains all initialized models used by the backend.
type Models struct {
	Product   core.ProductModel
	Purchaser core.PurchaserModel
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
	b.Models.Purchaser = mysql.NewPurchaserModel()
}

func (b *Backend) wireServices() {
	// Setup services.
	product := service.NewProductService()
	product.SetProductModel(b.Models.Product)

	purchaser := service.NewPurchaserService()
	purchaser.SetPurchaserModel(b.Models.Purchaser)

	b.Services.Product = product
	b.Services.Purchaser = purchaser

	util.EnsureNoNilPointers(
		b.Services.Product,
		b.Services.Purchaser,
	)
}

func (b *Backend) wireControllers() {
	// Setup controllers.
	con.ProductController.SetProductService(b.Services.Product)
	con.PurchaserController.SetPurchaserService(b.Services.Purchaser)

	util.EnsureNoNilPointers(
		con.ProductController,
		con.PurchaserController,
	)
}

// CreateTables creates the tables.
func (b *Backend) CreateTables() {
	ctx := context.Background()
	var err error
	panicOnError := func() {
		if err != nil {
			panic(err)
		}
	}

	err = b.Models.Product.CreateTable(ctx, true)
	panicOnError()

	err = b.Models.Purchaser.CreateTable(ctx, true)
	panicOnError()
}
