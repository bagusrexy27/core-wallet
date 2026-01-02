package models

import (
	"time"
)

type TransactionType string
type TransactionStatus string

const (
	TransactionTypeTopup    TransactionType = "TOPUP"
	TransactionTypeWithdraw TransactionType = "WITHDRAW"
	TransactionTypeTransfer TransactionType = "TRANSFER"

	TransactionStatusSuccess TransactionStatus = "SUCCESS"
	TransactionStatusFailed  TransactionStatus = "FAILED"
	TransactionStatusPending TransactionStatus = "PENDING"
)

type Transaction struct {
	ID            string            `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	WalletID      string            `json:"wallet_id" gorm:"type:uuid;not null"`
	Type          TransactionType   `json:"type" gorm:"type:varchar(20);not null"`
	Amount        int64             `json:"amount" gorm:"not null"`
	BalanceBefore int64             `json:"balance_before" gorm:"not null"`
	BalanceAfter  int64             `json:"balance_after" gorm:"not null"`
	PaymentMethod string            `json:"payment_method,omitempty" gorm:"size:50"`
	Status        TransactionStatus `json:"status" gorm:"type:varchar(20);not null"`
	CreatedAt     time.Time         `json:"created_at" gorm:"autoCreateTime"`
}

func (Transaction) TableName() string {
	return "transactions"
}
