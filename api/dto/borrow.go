package dto

type CreateBorrow struct {
	UserID string `json:"user_id,omitempty" form:"user_id"`
	BookID string `json:"book_id,omitempty" form:"book_id"`
}

type PutBorrow struct {
	ID     string `json:"id" form:"id" binding:"required"`
	UserID string `json:"user_id,omitempty" form:"user_id" binding:"required"`
	BookID string `json:"book_id,omitempty" form:"book_id" binding:"required"`
}

type PatchBorrow struct {
	ID     string `json:"id" form:"id" binding:"required"`
	UserID string `json:"user_id,omitempty" form:"user_id"`
	BookID string `json:"book_id,omitempty" form:"book_id"`
}
