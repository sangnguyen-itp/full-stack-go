package models

import (
	"time"
)

type Contract struct {
	ID        uint      `json:"id" gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`

	ContractTypeID uint         `json:"contract_type_id" `
	ContractType   ContractType `json:"contract_type" gorm:"foreignKey:ContractTypeID"`

	Value    float64 `json:"value"`
	Currency string  `json:"currency"`

	// Electric signature
	HashAlgorithm         string `json:"hash_algorithm"` // uni-directly
	HashedContractContent string `json:"hashed_content"`

	RandomPublicKey          string `json:"random_public_key"`
	RandomSecretKey          string `json:"random_secret_key"`
	CryptoAlgorithm          string `json:"crypto_algorithm"` // bi-directly
	EncryptedContractContent string `json:"encrypted_contract_content"`

	IsVisible       bool               `json:"is_visible"`
	Status          string             `json:"status"` // open|signed from customer|canceled|approved|rejected|closed
	StatusTimelines []ContractTimeline `json:"status_timelines" gorm:"foreignKey:ContractID"`

	// associations
	// belongs to
	CustomerId  uint     `json:"customer_id"`
	Customer    Customer `json:"customer"`
	CreatedById uint     `json:"created_by_id"`
	CreatedBy   User     `json:"created_by"`
}
