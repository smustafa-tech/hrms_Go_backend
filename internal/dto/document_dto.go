package dto

type DocumentUploadRequest struct {
	DocumentType string `form:"documentType" binding:"required"`
}
