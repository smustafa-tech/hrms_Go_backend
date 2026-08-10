package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Attendance struct {
	ID             uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID         string     `gorm:"not null;index" json:"userId"`
	Date           string     `gorm:"not null" json:"date"`
	Status         string     `gorm:"default:'present'" json:"status"`
	CheckIn        string     `json:"checkIn"`
	CheckOut       string     `json:"checkOut"`
	BreakStart     string     `json:"breakStart"`
	BreakEnd       string     `json:"breakEnd"`
	TotalBreakTime int        `gorm:"default:0" json:"totalBreakTime"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
}

func (a *Attendance) BeforeCreate(tx *gorm.DB) error {
	a.ID = uuid.New()
	return nil
}
