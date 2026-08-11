package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	SenderID     string    `gorm:"not null;index" json:"senderId"`
	ReceiverID   string    `gorm:"not null;index" json:"receiverId"`
	SenderType   string    `json:"senderType"`
	ReceiverType string    `json:"receiverType"`
	Type         string    `json:"type"`
	Message      string    `json:"message"`
	IsRead       bool      `gorm:"default:false" json:"isRead"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (n *Notification) BeforeCreate(tx *gorm.DB) error {
	n.ID = uuid.New()
	return nil
}
