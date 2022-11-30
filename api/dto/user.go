package dto

type CreateUser struct {
	Name  string `json:"name,omitempty" form:"name"`
	Email string `json:"email,omitempty" form:"email"`
	Phone string `json:"phone" form:"phone"`
}

type PutUser struct {
	ID    string `json:"id" form:"id" binding:"required"`
	Name  string `json:"name,omitempty" form:"name" binding:"required"`
	Email string `json:"email,omitempty" form:"email" binding:"required"`
	Phone string `json:"phone" form:"phone" binding:"required"`
}

type PatchUser struct {
	ID    string `json:"id" form:"id" binding:"required"`
	Name  string `json:"name,omitempty" form:"name"`
	Email string `json:"email,omitempty" form:"email"`
	Phone string `json:"phone" form:"phone"`
}
