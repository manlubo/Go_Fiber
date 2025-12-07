package entity

type Reply struct {
	ID      string  `json:"id"`
	Content *string `json:"content"`
	BoardID string  `json:"boardId"`
	UserID  string  `json:"userId"`
	BaseEntity
}
