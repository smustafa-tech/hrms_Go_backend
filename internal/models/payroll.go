package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payroll struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID      string     `gorm:"not null;index" json:"userId"`
	Month       int        `json:"month"`
	Year        int        `json:"year"`
	BasicSalary float64    `json:"basicSalary"`
	HRA         float64    `json:"hra"`
	Allowances  float64    `json:"allowances"`
	Deductions  float64    `json:"deductions"`
	NetSalary   float64    `json:"netSalary"`
	Status      string     `gorm:"default:'generated'" json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `gorm:"index" json:"-"`
}

func (p *Payroll) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}
