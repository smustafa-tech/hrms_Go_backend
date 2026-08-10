package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName        string    `gorm:"not null" json:"firstName"`
	LastName         string    `gorm:"not null" json:"lastName"`
	Email            string    `gorm:"uniqueIndex;not null" json:"email"`
	Password         string    `gorm:"not null" json:"-"`
	OrganizationName string    `json:"organizationName"`
	Phone            string    `json:"phone"`
	Slug             string    `gorm:"uniqueIndex;not null" json:"slug"`
	Role             string    `gorm:"default:'admin'" json:"role"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
