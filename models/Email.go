package models

import (
	"time"
)

type Email struct {
	ID        uint64    `json:"id" gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`

	Value string `json:"value"`
	IsPrimary bool `json:"is_primary"`
	
	OwnerId uint `json:"owner_id"`
}
