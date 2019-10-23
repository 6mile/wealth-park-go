package mysql

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/yashmurty/wealth-park/wpark/core"
	"go.uber.org/zap"
)

// ProductModel ...
type ProductModel struct {
	BasicModel
}

var (
// _ core.ProductModel = &ProductModel{}
)

// NewProductModel ...
func NewProductModel() *ProductModel {
	return &ProductModel{
		BasicModel: BasicModel{
			tableName: "wpark_product",
			tableSchema: `CREATE TABLE IF NOT EXISTS wpark_product (
			id VARCHAR(36) NOT NULL UNIQUE,
			created_at BIGINT NOT NULL,
			updated_at BIGINT NOT NULL,
			is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
			name VARCHAR(255) NOT NULL,
			PRIMARY KEY ( id )
		) ENGINE=InnoDB`,
		},
	}
}

// Create ...
func (s *ProductModel) Create(ctx context.Context, d *core.Product) error {
	method := "create"

	defer logCall(ctx, time.Now().UTC(), mysqlTag, method, "",
		zap.String("table", s.tableName),
		zap.String("id", d.ID),
	)

	prepString := `INSERT INTO ` + s.tableName + `
	(
		id,
		created_at,
		updated_at,
		is_deleted,
		name
	)
	VALUES (
		?, ?, ?, ?, ?
	)`

	stmt, err := db.PrepareContext(ctx, prepString)
	if err != nil {
		log.Error(mysqlTag+": error",
			zap.Error(err),
			zap.String("table", s.tableName),
			zap.String("method", method),
			zap.String("id", d.ID),
		)
		return errors.Wrapf(err, mysqlTag+": "+method+" failed for %s in %s. Could not prepare db statement", d.ID, s.tableName)
	}
	res, err := stmt.ExecContext(ctx,
		d.ID,
		d.CreatedAt,
		d.UpdatedAt,
		d.IsDeleted,
		d.Name,
	)
	if err != nil {
		log.Error(mysqlTag+": error",
			zap.Error(err),
			zap.String("table", s.tableName),
			zap.String("method", method),
			zap.String("id", d.ID),
		)
		return errors.Wrapf(err, mysqlTag+": "+method+" failed for %s in %s. Could not execute db statement", d.ID, s.tableName)
	}
	// lastId returned will be zero when if pass ID manually in the payload.
	_, err = res.LastInsertId()
	if err != nil {
		log.Error(mysqlTag+": error",
			zap.Error(err),
			zap.String("table", s.tableName),
			zap.String("method", method),
			zap.String("id", d.ID),
		)
		return errors.Wrapf(err, mysqlTag+": "+method+" failed for %s in %s. Could not get last insert id", d.ID, s.tableName)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Error(mysqlTag+": error",
			zap.Error(err),
			zap.String("table", s.tableName),
			zap.String("method", method),
			zap.String("id", d.ID),
		)
		return errors.Wrapf(err, mysqlTag+": "+method+" failed for %s in %s. Could not get affected rows", d.ID, s.tableName)
	}
	// Throw error if row count is zero.
	if rowCnt <= 0 {
		err := fmt.Errorf("query returned empty row count")
		log.Error(mysqlTag+": error",
			zap.Error(err),
			zap.String("table", s.tableName),
			zap.String("method", method),
			zap.String("id", d.ID),
		)
		return errors.Wrapf(err, mysqlTag+": "+method+" failed for %s in %s. Affected row count should be greater than 0", d.ID, s.tableName)
	}

	return nil
}
