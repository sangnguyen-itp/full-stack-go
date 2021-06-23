package models

import (
	"time"
)

type Policy struct {
	ID        uint    `json:"id" gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`

	Name string `json:"name"`
	Type string `json:"type"`
	
	//associations
	// belongs to
	BankId uint `json:"bank_id"`
	ContractTypeId uint `json:"-"`
	
	// has many
	InterestRates []InterestRate `json:"interest_rates"`
}