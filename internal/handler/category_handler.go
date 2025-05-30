package handler

import (
	"company-profile/internal/dto/request"
	"company-profile/internal/dto/response"
	"company-profile/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CategoryHandler struct {
	service usecase.CategoryUsecase
}

func NewCategoryHandler(s usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{s}
}

// GetAllCategories godoc
// @Summary Get all categories
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

// GetCategoryByID godoc
// @Summary Get category by ID
// @Tags Category
// @Security BearerAuth
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} response.CategoryResponse
// @Failure 404 {object} response.MessageResponse
// @Router /api/category/{id} [get]
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, response.MessageResponse{
			Message: "Data tidak ditemukan!",
		})
		return
	}
	c.JSON(http.StatusOK, response.MessageResponse{
		Message: "success",
		Data:    category,
	})
}

// UpdateCategory godoc
// @Summary Update category
// @Tags Category
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param data body request.UpdateCategoryRequest true "Category data"
// @Success 200 {object} response.CategoryResponse
// @Failure 400 {object} response.MessageResponse
// @Failure 404 {object} response.MessageResponse
// @Router /api/category/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req request.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.MessageResponse{
			Message: "Data tidak valid!",
		})
		return
	}

	category, err := h.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, response.MessageResponse{
			Message: "Data tidak ditemukan!",
		})
		return
	}

	c.JSON(http.StatusOK, response.MessageResponse{
		Message: "success",
		Data:    category,
	})
}

// DeleteCategory godoc
// @Summary Delete category
// @Tags Category
// @Security BearerAuth
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} response.MessageResponse
// @Failure 404 {object} response.MessageResponse
// @Router /api/category/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, response.MessageResponse{
			Message: "Data tidak di temukan!",
		})
		return
	}

	c.JSON(http.StatusOK, response.MessageResponse{
		Message: "Data berhasil di hapus!",
	})
}
