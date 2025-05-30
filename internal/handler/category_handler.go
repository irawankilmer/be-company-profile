package handler

import (
	"company-profile/internal/dto/request"
	"company-profile/internal/dto/response"
	"company-profile/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CategoryHandler struct {
	service usecase.CategoryUsecase
}

func NewCategoryHandler(s usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{s}
}

// GetAllCategory godoc
// @Summary Get all posts
// @Tags Category
// @Security BearerAuth
// @Produce json
// @Success 200 {array} response.CategoryResponse
// @Failure 401 {object} response.MessageResponse
// @Router /api/category [get]
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengambil data kategori",
		})
		return
	}

	c.JSON(http.StatusOK, response.MessageResponse{
		Message: "success",
		Data:    categories,
	})
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
