package handler

import (
	"company-profile/internal/dto/request"
	"company-profile/internal/dto/response"
	"company-profile/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type PostHandler struct {
	service service.PostService
}

func NewPostHandler(s service.PostService) *PostHandler {
	return &PostHandler{s}
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
