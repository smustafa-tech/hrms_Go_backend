package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	EmpID         string     `gorm:"uniqueIndex;not null" json:"emp_id"`
	FirstName     string     `gorm:"not null" json:"firstName"`
	MiddleName    string     `json:"middleName"`
	LastName      string     `gorm:"not null" json:"lastName"`
	Email         string     `gorm:"uniqueIndex;not null" json:"email"`
	Phone         string     `json:"phone"`
	AdharCard     string     `json:"adharCard"`
	Designation   string     `json:"designation"`
	Role          string     `gorm:"default:'employee'" json:"role"`
	Department    string     `json:"department"`
	DateOfJoining string     `json:"dateOfJoining"`
	Status        string     `gorm:"default:'active'" json:"status"`
	Salary        float64    `json:"salary"`
	WorkMode      string     `gorm:"default:'office'" json:"workMode"`
	MgrID         string     `json:"mgrId"`
	HrID          string     `json:"hrId"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `gorm:"index" json:"-"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) error {
	e.ID = uuid.New()
	return nil
}
