package handler

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smustafa-tech/hrms-backend/internal/dto"
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"github.com/smustafa-tech/hrms-backend/internal/service"
)

type DocumentHandler struct {
	svc *service.DocumentService
}

func NewDocumentHandler(svc *service.DocumentService) *DocumentHandler {
	return &DocumentHandler{svc: svc}
}

func (h *DocumentHandler) UploadDocument(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req dto.DocumentUploadRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "file is required"})
		return
	}

	uploadDir := filepath.Join("uploads", "documents")
	if err := os.MkdirAll(uploadDir, 0o755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create upload directory"})
		return
	}

	fileName := uuid.New().String() + filepath.Ext(file.Filename)
	filePath := filepath.Join(uploadDir, fileName)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save document"})
		return
	}

	doc := &models.Document{
		UserID:       userID.(string),
		DocumentType: req.DocumentType,
		FileName:     file.Filename,
		FilePath:     filePath,
		MimeType:     file.Header.Get("Content-Type"),
		Size:         int(file.Size),
	}

	if err := h.svc.CreateDocument(doc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"document": doc})
}

func (h *DocumentHandler) GetMyDocuments(c *gin.Context) {
	userID, _ := c.Get("userID")
	docs, err := h.svc.GetDocumentsForUser(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"documents": docs})
}

func (h *DocumentHandler) GetEmployeeDocuments(c *gin.Context) {
	role, _ := c.Get("role")
	allowed := role.(string) == "admin" || role.(string) == "hr"
	if !allowed {
		c.JSON(http.StatusForbidden, gin.H{"message": "forbidden"})
		return
	}

	docs, err := h.svc.GetAllDocuments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"documents": docs})
}

func (h *DocumentHandler) DownloadDocument(c *gin.Context) {
	docID := c.Param("id")
	doc, err := h.svc.FindDocumentByID(docID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "document not found"})
		return
	}

	if _, err := os.Stat(doc.FilePath); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "file not found"})
		return
	}
	c.FileAttachment(doc.FilePath, doc.FileName)
}

func (h *DocumentHandler) DeleteDocument(c *gin.Context) {
	docID := c.Param("id")
	doc, err := h.svc.FindDocumentByID(docID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "document not found"})
		return
	}

	if err := h.svc.DeleteDocument(docID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := os.Remove(doc.FilePath); err != nil && !os.IsNotExist(err) {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to remove file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "document deleted"})
}
