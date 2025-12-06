package entity

type Board struct {
	ID      string  `json:"id"`
	UserID  string  `json:"userId"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
}
