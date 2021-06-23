package models

import (
	"time"
)

type InterestRate struct {
	ID        uint64    `json:"id" gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	
	Duration int `json:"duration"`
	DurationType time.Duration `json:"duration_type"` // Month, year
	Value float32 `json:"value"`
	
	//associations
	// belongs to
	PolicyId uint `json:"policy_id"`
	Policy Policy `json:"policy"`
}
