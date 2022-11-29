package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Borrow struct {
	ID        string         `json:"id" gorm:"type:varchar(36);primaryKey;not null"`
	UserID    string         `json:"user_id" gorm:"type:varchar(36);not null"`
	BookID    string         `json:"book_id" gorm:"type:varchar(36);not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Users     User           `json:"users" gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Books     Book           `json:"books" gorm:"foreignKey:BookID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
}

func (u *Borrow) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()

	return
}
