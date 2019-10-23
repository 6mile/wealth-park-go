package core

// Resource represents a basic resource (promotion, banner, etc).
type Resource struct {
	ID        string `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	IsDeleted bool   `json:"is_deleted"`
}
