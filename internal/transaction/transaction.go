package transaction

import (
	"time"
)

type Transaction struct {
	ID           int              `gorm:"primaryKey"`
	CustomerName string           `gorm:"size:255;not null"`
	TotalPoints  int              `gorm:"not null"`
	Items        []TransactionItem `gorm:"foreignKey:TransactionID"` 
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type TransactionItem struct {
	ID            int      `gorm:"primaryKey"`
	TransactionID int      `gorm:"not null"`
	VoucherID     int      `gorm:"not null"`
	Quantity      int      `gorm:"not null"`
	TotalPoints   int      `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type RedemptionRequest struct {
	CustomerName string `json:"customer_name"`
	VoucherItems []struct {
		VoucherID int `json:"voucher_id"`
		Quantity  int `json:"quantity"`
		Points    int `json:"points"`
	} `json:"voucher_items"`
}
