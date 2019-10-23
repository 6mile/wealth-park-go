package mysql

import (
	"context"
)

// BasicModel contains basic operations.
type BasicModel struct {
	tableName   string
	tableSchema string
}

// CreateTable creates a new table using name and schema.
func (m *BasicModel) CreateTable(ctx context.Context, delete bool) error {
	return CreateTable(ctx, m.tableName, m.tableSchema, true)
}
