package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QueryReply struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	QueryID   string     `gorm:"not null;index" json:"queryId"`
	RepliedBy string     `gorm:"not null" json:"repliedBy"`
	Message   string     `json:"message"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
}

func (qr *QueryReply) BeforeCreate(tx *gorm.DB) error {
	qr.ID = uuid.New()
	return nil
}
