package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Document struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID       string     `gorm:"not null;index" json:"userId"`
	DocumentType string     `json:"documentType"`
	FileName     string     `json:"fileName"`
	FilePath     string     `json:"filePath"`
	MimeType     string     `json:"mimeType"`
	Size         int        `json:"size"`
	UploadedBy   string     `json:"uploadedBy"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `gorm:"index" json:"-"`
}

func (d *Document) BeforeCreate(tx *gorm.DB) error {
	d.ID = uuid.New()
	return nil
}
