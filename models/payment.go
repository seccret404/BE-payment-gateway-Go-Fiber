package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	OrderID       string         `gorm:"uniqueIndex" json:"order_id"`
	ProductID     string         `json:"product_id"`
	Quantity      int            `json:"quantity"`       
	CustomerName  string         `json:"customer_name"`
	CustomerEmail string         `json:"customer_email"`
	Amount        int64          `json:"amount"`
	Status        string         `json:"status"`         
	SnapURL       string         `json:"snap_url"`
	PaymentType   string         `json:"payment_type"`   
	PaidAt        *time.Time     `json:"paid_at"`        	
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}