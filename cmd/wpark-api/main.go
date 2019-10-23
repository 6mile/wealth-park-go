package main

import (
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yashmurty/wealth-park/wpark/apiserver"
	"github.com/yashmurty/wealth-park/wpark/backend"
	"github.com/yashmurty/wealth-park/wpark/config"
	"github.com/yashmurty/wealth-park/wpark/controller"
)

var (
	InitDB              bool
	PreCommitCheck      bool
	NoBackgroundThreads bool
)

func main() {
	flag.BoolVar(&InitDB, "initdb", false, "Drop the database and recreate all indexes")
	flag.BoolVar(&PreCommitCheck, "pre-commit", false, "Perform a pre-commit check and exit")
	flag.Parse()

	// Load config.
	c := config.GetInstance()

	// Create an MySQL based backend.
	be := backend.NewBackendWithMYSQLModels()

	// Check if we should (re)initialize DB.
	if InitDB {
		err := initDB(be)
		if err != nil {
			panic(err)
		} else {
			println("It's a Done Deal.")
			os.Exit(0)
		}
	}

	// Setup API gateway.
	if c.IsProduction {
		// Set Gin release mode in production.
		gin.SetMode(gin.ReleaseMode)
	}

	// Create an API server instance.
	be.Server = apiserver.NewServer(apiserver.NewServerArgs{
		Addr: c.GetAddr(),
	})

	controller.SetupHTTPHandlers(be.Server)

	// Exit at this point if this is a simple pre-commit check.
	if PreCommitCheck {
		println("All is good in the hood!")
		os.Exit(0)
	}
	be.Server.Start()
}

// initDB deletes all existing tables and recreates them.
func initDB(be *backend.Backend) error {
	be.CreateTables()
	return nil
}
