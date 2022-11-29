package dto

type CreateBook struct {
	Name string `json:"name,omitempty" form:"name"`
	Page string `json:"page,omitempty" form:"page"`
}

type UpdateBook struct {
	ID   string `json:"id" form:"id"`
	Name string `json:"name,omitempty" form:"name"`
	Page string `json:"page,omitempty" form:"page"`
}
