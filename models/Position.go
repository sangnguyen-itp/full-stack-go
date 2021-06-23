package models

import (
	"time"
)

type Position struct {
	ID        uint64    `json:"id" gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`

	Name string `json:"name"`
	
	//associations
	// belongs to
	BankId uint `json:"-"`
	Bank Bank `json:"bank"`

	// has many
	Users []User `json:"-" gorm:"foreignKey:PositionId"`
}