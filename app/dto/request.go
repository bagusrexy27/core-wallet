package dto

type CreateWalletRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

type TopUpRequest struct {
	WalletID      string `json:"wallet_id" validate:"required"`
	Amount        int64  `json:"amount" validate:"required,gt=0"`
	PaymentMethod string `json:"payment_method" validate:"required"`
}

type ConfirmTopUpRequest struct {
	WalletID      string `json:"wallet_id" validate:"required"`
	TransactionID string `json:"transaction_id" validate:"required"`
}

type CheckStatusTopUpRequest struct {
	TransactionID string `json:"transaction_id" validate:"required"`
}
