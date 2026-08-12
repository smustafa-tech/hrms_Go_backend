package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Query struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID     string     `gorm:"not null;index" json:"userId"`
	Subject    string     `json:"subject"`
	Message    string     `json:"message"`
	Priority   string     `json:"priority"`
	Status     string     `gorm:"default:'open'" json:"status"`
	Attachment string     `json:"attachment"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `gorm:"index" json:"-"`
}

func (q *Query) BeforeCreate(tx *gorm.DB) error {
	q.ID = uuid.New()
	return nil
}
