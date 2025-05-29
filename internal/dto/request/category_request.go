package request

type CreateCategoryRequest struct {
	Name             string `json:"name" binding:"required"`
	Description      string `json:"description" binding:"required"`
	ParentCategoryID *uint  `json:"parent_category_id" binding:"omitempty"`
}

type UpdateCategoryRequest struct {
	Name             string `json:"name" binding:"required"`
	Description      string `json:"description" binding:"required"`
	ParentCategoryID *uint  `json:"parent_category_id" binding:"omitempty"`
}
