package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentLog struct {
	gorm.Model
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	TransactionID uuid.UUID      `gorm:"type:uuid" json:"transaction_id"`
	Status        string         `json:"status"`
	RawResponse   string         `gorm:"type:text" json:"raw_response"`
	CreatedAt     time.Time      `json:"created_at"`

	Payment   Payment     `gorm:"foreignKey:TransactionID"`
}