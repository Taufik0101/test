package dto

type CreateUser struct {
	Name  string `json:"name,omitempty" form:"name"`
	Email string `json:"email,omitempty" form:"email"`
	Phone uint64 `json:"phone" form:"phone"`
}

type UpdateUser struct {
	ID    string `json:"id" form:"id"`
	Name  string `json:"name,omitempty" form:"name"`
	Email string `json:"email,omitempty" form:"email"`
	Phone uint64 `json:"phone" form:"phone"`
}
