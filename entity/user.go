package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key;"`
	UserName    string    `json:"userName" gorm:"type:varchar(255);not null;" binding:"required"`
	Email       string    `json:"email" gorm:"type:varchar(255);not null;unique" binding:"required"`
	Password    string    `json:"password" gorm:"type:varchar(255);not null;" binding:"required"`
	Role        uint      `json:"role"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
