package entity

type BaseEntity struct {
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
	IsActive  bool   `json:"isActive"`
	IsDeleted bool   `json:"isDeleted"`
	DeletedAt *int64 `json:"deletedAt"`
}
