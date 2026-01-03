package models

import "time"

type Wallet struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserId    string    `json:"user_id" gorm:"index"`
	Balance   int64     `json:"balance"`
	Checksum  string    `json:"checksum"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (w *Wallet) TableName() string {
	return "wallet"
}
