package models

import (
	"time"
)

type Customer struct {
	ID        uint    `json:"id" gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`

	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Emails []Email `json:"emails" gorm:"foreignKey:OwnerId"`

	Username string `json:"username"`
	Password string `json:"password"`
	IsSetTokenTimeout bool `json:"account_timeout"`
	TimeoutDuration int `json:"timeout_duration"` //duration minute

	// associations
	// many to many
	Contracts []Contract `json:"contracts" gorm:"foreignKey:CustomerId;"`
}