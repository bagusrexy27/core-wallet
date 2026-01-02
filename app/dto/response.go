package dto

import "time"

type UserLoginResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateWalletResponse struct {
	WalletID string `json:"wallet_id"`
	UserID   string `json:"user_id"`
	Balance  int64  `json:"balance"`
}

type BalanceResponse struct {
	WalletId    string    `json:"wallet_id"`
	Balance     int64     `json:"balance"`
	LastUpdated time.Time `json:"last_updated"`
}
