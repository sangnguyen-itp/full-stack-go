package models

import (
	"time"
)

type ContractTimeline struct {
	ID        uint    `json:"id" gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`

	Description string `json:"description"`
	ContractID uint   `json:"created_by_id"`
}
