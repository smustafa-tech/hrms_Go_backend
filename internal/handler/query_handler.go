package handler

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"github.com/smustafa-tech/hrms-backend/internal/service"
)

type QueryHandler struct {
	svc *service.QueryService
}

func NewQueryHandler(svc *service.QueryService) *QueryHandler {
	return &QueryHandler{svc: svc}
}

func (h *QueryHandler) GetQueriesUnreadCount(c *gin.Context) {
	userID, _ := c.Get("userID")
	count, err := h.svc.GetUnreadCount(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

func (h *QueryHandler) GetMyQueries(c *gin.Context) {
	userID, _ := c.Get("userID")
	queries, err := h.svc.GetMyQueries(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, queries)
}

func (h *QueryHandler) GetAllQueries(c *gin.Context) {
	queries, err := h.svc.GetAllQueries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, queries)
}

func (h *QueryHandler) SubmitQuery(c *gin.Context) {
	subject := c.PostForm("subject")
	message := c.PostForm("message")
	priority := c.PostForm("priority")
	attachmentPath := ""

	file, err := c.FormFile("attachment")
	if err == nil {
		uploadsDir := filepath.Join("uploads", "queries")
		if err := os.MkdirAll(uploadsDir, 0o755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create upload directory"})
			return
		}

		fileName := uuid.New().String() + filepath.Ext(file.Filename)
		filePath := filepath.Join(uploadsDir, fileName)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save attachment"})
			return
		}
		attachmentPath = filePath
	}

	userID, _ := c.Get("userID")
	query := &models.Query{
		UserID:     userID.(string),
		Subject:    subject,
		Message:    message,
		Priority:   priority,
		Status:     "open",
		Attachment: attachmentPath,
	}

	newQuery, err := h.svc.SubmitQuery(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newQuery)
}

func (h *QueryHandler) ReplyToQuery(c *gin.Context) {
	queryID := c.Param("id")
	var payload struct {
		Message string `json:"message" binding:"required"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID, _ := c.Get("userID")
	reply := &models.QueryReply{
		QueryID:   queryID,
		RepliedBy: userID.(string),
		Message:   payload.Message,
	}

	created, err := h.svc.ReplyToQuery(reply)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, created)
}

func (h *QueryHandler) CloseQuery(c *gin.Context) {
	queryID := c.Param("id")
	if err := h.svc.CloseQuery(queryID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "query closed"})
}

func (h *QueryHandler) DeleteQuery(c *gin.Context) {
	queryID := c.Param("id")
	if err := h.svc.DeleteQuery(queryID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "query deleted"})
}
