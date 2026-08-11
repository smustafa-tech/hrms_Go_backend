package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KraKpi struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID     string     `gorm:"not null;index" json:"userId"`
	Title      string     `json:"title"`
	Description string    `json:"description"`
	KRA        string     `json:"kra"`
	KPI        string     `json:"kpi"`
	Target     float64    `json:"target"`
	Achieved   float64    `json:"achieved"`
	Rating     int        `json:"rating"`
	Comments   string     `json:"comments"`
	Month      int        `json:"month"`
	Year       int        `json:"year"`
	Status     string     `gorm:"default:'PENDING'" json:"status"`
	ManagerID  string     `json:"managerId"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `gorm:"index" json:"-"`
}

func (k *KraKpi) BeforeCreate(tx *gorm.DB) error {
	k.ID = uuid.New()
	return nil
}
