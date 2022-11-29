package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID        string         `json:"id" gorm:"type:varchar(36);primaryKey;not null"`
	Name      string         `json:"name" gorm:"type:varchar(255);not null"`
	Page      string         `json:"page" gorm:"type:varchar(20)"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (u *Book) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()

	return
}
