package dto

type CreateBook struct {
	Name string `json:"name,omitempty" form:"name"`
	Page string `json:"page,omitempty" form:"page"`
}

type PutBook struct {
	ID   string `json:"id" form:"id" binding:"required"`
	Name string `json:"name,omitempty" form:"name" binding:"required"`
	Page string `json:"page,omitempty" form:"page" binding:"required"`
}

type PatchBook struct {
	ID   string `json:"id" form:"id" binding:"required"`
	Name string `json:"name,omitempty" form:"name"`
	Page string `json:"page,omitempty" form:"page"`
}
