package core

import "context"

// Model describes a set of common data layer operations.
type Model interface {
	// CreateTable recreates the index.
	CreateTable(ctx context.Context, delete bool) error
}
