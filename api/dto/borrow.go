package dto

type CreateBorrow struct {
	UserID string `json:"user_id,omitempty" form:"user_id"`
	BookID string `json:"book_id,omitempty" form:"book_id"`
}

type UpdateBorrow struct {
	ID     string `json:"id" form:"id"`
	UserID string `json:"user_id,omitempty" form:"user_id"`
	BookID string `json:"book_id,omitempty" form:"book_id"`
}
