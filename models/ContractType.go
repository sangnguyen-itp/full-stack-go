package models

import (
	"time"
)

type ContractType struct {
	ID        uint    `json:"id" gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	
	Name string `json:"name"`
	
	// associations
	// belongs to
	BankId uint `json:"bank_id"`
	Bank Bank `json:"bank"`

	// has one
	Policy Policy `json:"policy"`
}