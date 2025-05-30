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

type PostHandler struct {
	service usecase.PostUsecase
}

func NewPostHandler(s usecase.PostUsecase) *PostHandler {
	return &PostHandler{s}
}

// GetAllPosts godoc
// @Summary Get all posts
// @Tags Post
// @Security BearerAuth
// @Produce json
// @Success 200 {array} response.PostResponse
// @Failure 401 {object} response.MessageResponse
// @Router /api/post [get]
func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengambil data post",
		})
		return
	}
	c.JSON(http.StatusOK, response.MessageResponse{
		Message: "success",
		Data:    posts,
	})
}

// CreatePost godoc
// @Security BearerAuth
// @Summary Create new post
// @Tags Post
// @Accept json
// @Produce json
// @Param request body request.CreatePostRequest true "Post data"
// @Success 201 {object} response.PostResponse
// @Failure 400 {object} response.MessageResponse
// @Router /api/post [post]
func (h *PostHandler) CreatePost(c *gin.Context) {
	var req request.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.MessageResponse{
			Message: "Input tidak sesuai",
		})
		return
	}

	userID := c.GetUint("user_id")
	post, err := h.service.Create(req, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.MessageResponse{
			Message: "Gagal membuat postingan!",
		})
		return
	}

	c.JSON(http.StatusOK, response.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		CreatedAt: post.CreatedAt.Format(time.RFC3339),
		UpdatedAt: post.UpdatedAt.Format(time.RFC3339),
	})
}

// GetPostByID godoc
// @Summary Get post by ID
// @Tags Post
// @Security BearerAuth
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} response.PostResponse
// @Failure 404 {object} response.MessageResponse
// @Router /api/post/{id} [get]
func (h *PostHandler) GetPostByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data tidak ditemukan!",
		})
		return
	}
	c.JSON(http.StatusOK, response.MessageResponse{
		Message: "success",
		Data:    post,
	})
}

// UpdatePost godoc
// @Summary Update post
// @Tags Post
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param data body request.UpdatePostRequest true "Post data"
// @Success 200 {object} response.PostResponse
// @Failure 400 {object} response.MessageResponse
// @Failure 404 {object} response.MessageResponse
// @Router /api/post/{id} [put]
func (h *PostHandler) UpdatePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req request.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.MessageResponse{
			Message: "Data tidak valid!",
		})
		return
	}

	post, err := h.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, response.MessageResponse{
			Message: "Data tidak ditemukan!",
		})
		return
	}

	c.JSON(http.StatusOK, response.MessageResponse{
		Message: "success",
		Data:    post,
	})
}

// DeletePost godoc
// @Summary Delete post
// @Tags Post
// @Security BearerAuth
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} response.MessageResponse
// @Failure 404 {object} response.MessageResponse
// @Router /api/post/{id} [delete]
func (h *PostHandler) DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, response.MessageResponse{
			Message: "Postingan tidak ditemukan!",
		})
		return
	}
	c.JSON(http.StatusOK, response.MessageResponse{
		Message: "Postingan berhasil di hapus!",
	})
}
