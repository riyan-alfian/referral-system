package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Email        string    `gorm:"unique;not null" json:"email" binding:"required,email"`
	Password     string    `gorm:"not null" json:"-"`
	ReferralLink string    `gorm:"unique" json:"referral_link"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Contribution struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Email        string    `gorm:"unique;not null" json:"email" binding:"required,email"`
	ReferralLink string    `json:"referral_link"`
	CreatedAt    time.Time `json:"created_at"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()
	return scope.SetColumn("ID", uuid)
}

func (contribution *Contribution) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()
	return scope.SetColumn("ID", uuid)
}
