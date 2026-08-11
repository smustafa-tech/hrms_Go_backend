package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Leave struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID      string     `gorm:"not null;index" json:"userId"`
	LeaveType   string     `json:"leaveType"`
	StartDate   string     `json:"startDate"`
	EndDate     string     `json:"endDate"`
	TotalDays   int        `json:"totalDays"`
	Reason      string     `json:"reason"`
	Status      string     `gorm:"default:'PENDING'" json:"status"`
	ApprovedBy  string     `json:"approvedBy"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `gorm:"index" json:"-"`
}

func (l *Leave) BeforeCreate(tx *gorm.DB) error {
	l.ID = uuid.New()
	return nil
}
