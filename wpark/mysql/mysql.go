package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/yashmurty/wealth-park/wpark/config"
	"github.com/yashmurty/wealth-park/wpark/pkg/logger"
	"go.uber.org/zap"

	// Blank import to pass the mysql driver to database/sql Open function
	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlTag = "mysql"
	db       *sql.DB
	log      = logger.Get(mysqlTag)
	// mysqlURL fetched from config
	clientErrorMode = false
)

// SetupDBHandle : We should understand the concept that a sql.DB is not a connection.
// When you use sql.Open() you get a handle for a database.
func SetupDBHandle() *sql.DB {
	if clientErrorMode {
		// Simulate connection failure.
		db = nil
		panic("clientErrorMode set to true")
	}

	if db != nil {
		return db
	}
	c := config.GetInstance()

	var err error

	db, err = sql.Open("mysql", c.MySQLURL+"/"+c.MySQLDBName)
	if err != nil {
		log.With(
			zap.Error(err),
			zap.String("mysql_url", c.MySQLURL),
		).Panic("could not connect to mysql")
	}

	return db
}

// PingServer will ping the MySQL server.
func PingServer(ctx context.Context) (err error) {
	method := "ping server"

	c := config.GetInstance()

	err = db.Ping()
	if err != nil {
		log.Error(mysqlTag+": error",
			zap.Error(err),
			zap.String("mysql_url", c.MySQLURL),
		)
		return errors.Wrapf(err, mysqlTag+": "+method+" ping failed")
	}

	log.Info(mysqlTag+": "+method+" successfull. Connected to mysql",
		zap.String("mysql_url", c.MySQLURL),
	)

	return
}

// CreateTable creates a MySQL table.
func CreateTable(ctx context.Context, tableName, tableSchema string, delete bool) error {
	method := "create table"

	defer logCall(ctx, time.Now(), mysqlTag, method, "", zap.String("table", tableName))

	// Disable foreign key constraint check.
	err := setForeignKeyCheck(tableName, false)
	if err != nil {
		return errors.Wrapf(err, mysqlTag+": "+method+" failed. Could not disable foreign key check")
	}

	// Check if table already exists.
	exists := checkTableExists(tableName)

	deleted := false
	if exists {
		if delete {
			err = deleteTable(tableName)
			if err != nil {
				return errors.Wrapf(err, mysqlTag+": "+method+" failed. Could not drop table")
			}
			deleted = true
		}
	}

	if !exists || deleted {
		// Create a new table.
		_, err = db.Exec(tableSchema)
		if err != nil {
			log.Error(mysqlTag+": error",
				zap.String("table", tableName),
				zap.Error(err),
			)
			return errors.Wrapf(err, mysqlTag+": "+method+" failed. Could not create table")
		}
	}

	// Enable foreign key constraint check.
	err = setForeignKeyCheck(tableName, true)
	if err != nil {
		return errors.Wrapf(err, mysqlTag+": "+method+" failed. Could not enable foreign key check")
	}

	return nil
}

func setForeignKeyCheck(tableName string, enable bool) error {
	checkValue := 0
	if enable {
		checkValue = 1
	}
	execString := "SET FOREIGN_KEY_CHECKS=" + strconv.Itoa(checkValue) + ";"
	_, err := db.Exec(execString)
	if err != nil {
		log.Error(mysqlTag+": error",
			zap.String("table", tableName),
			zap.Error(err),
		)
		return err
	}
	return nil
}

func checkTableExists(tableName string) bool {
	_, err := db.Exec(fmt.Sprintf("SELECT 1 FROM %s LIMIT 1;", tableName))
	if err != nil {
		return false
	}
	return true
}

func deleteTable(tableName string) error {
	_, err := db.Exec(fmt.Sprintf("DROP TABLE %s;", tableName))
	if err != nil {
		log.Error(mysqlTag+": error",
			zap.String("table", tableName),
			zap.Error(err),
		)
		return err
	}
	return nil
}

func logCall(ctx context.Context, start time.Time, tag, where, msg string, args ...zap.Field) {
	args = append(args, zap.String("took", time.Since(start).String()))
	args = append(args, zap.String("tag", tag), zap.String("where", where))
	log.Info(msg, args...)
}
