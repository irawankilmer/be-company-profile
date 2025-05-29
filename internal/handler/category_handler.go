package handler

import (
	"company-profile/internal/dto/request"
	"company-profile/internal/dto/response"
	"company-profile/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(s service.CategoryService) *CategoryHandler {
	return &CategoryHandler{s}
}

// CreateCategory godoc
// @Security BearerAuth
// @Summary Create new category
// @Tags Category
// @Accept json
// @Produce json
// @Param request body request.CreateCategoryRequest true "Category data"
// @Success 201 {object} response.CategoryResponse
// @Failure 400 {object} response.MessageResponse
// @Router /api/category [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req request.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.MessageResponse{
			Message: "Input tidak sesuai",
		})
		return
	}

	category, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.MessageResponse{
			Message: "Gagal membuat category!",
		})
		return
	}

	c.JSON(http.StatusOK, response.CategoryResponse{
		ID:               category.ID,
		Name:             category.Name,
		Description:      category.Description,
		ParentCategoryID: category.ParentCategoryID,
		CreatedAt:        category.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        category.UpdatedAt.Format(time.RFC3339),
	})
}
