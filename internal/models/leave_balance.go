package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LeaveBalance struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	EmployeeID   string    `gorm:"not null;index" json:"employeeId"`
	Year         int       `json:"year"`
	Casual       int       `gorm:"default:12" json:"casual"`
	Sick         int       `gorm:"default:8" json:"sick"`
	Earned       int       `gorm:"default:15" json:"earned"`
	Optional     int       `gorm:"default:2" json:"optional"`
	UsedCasual   int       `gorm:"default:0" json:"usedCasual"`
	UsedSick     int       `gorm:"default:0" json:"usedSick"`
	UsedEarned   int       `gorm:"default:0" json:"usedEarned"`
	UsedOptional int       `gorm:"default:0" json:"usedOptional"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (lb *LeaveBalance) BeforeCreate(tx *gorm.DB) error {
	lb.ID = uuid.New()
	return nil
}
