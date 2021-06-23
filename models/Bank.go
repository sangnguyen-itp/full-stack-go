package models

import "time"

type Bank struct {
	ID        uint    `json:"id" gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`

	Name string `json:"name"`

	// associations
	// has many
	Positions []Position `json:"positions" gorm:"foreignKey:BankId"`
	Users     []User     `json:"users" gorm:"foreignKey:BankId"`
	Policies  []Policy   `json:"policies" gorm:"foreignKey:BankId"`
}
