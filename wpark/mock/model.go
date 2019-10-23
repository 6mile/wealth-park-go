package mock

import (
	"context"

	"github.com/yashmurty/wealth-park/wpark/core"
)

// BasicModel ...
type BasicModel struct {
	CreateTableFn       func(ctx context.Context, delete bool) error
	CreateTableFnCalled int
}

var (
	_ core.Model = &BasicModel{}
)

// CreateTable ...
func (m *BasicModel) CreateTable(ctx context.Context, delete bool) error {
	return nil
}
