package handler

import "github.com/core-wallet/app/repository"

type WithdrawalHandler struct {
	transactionRepo *repository.TransactionRepo
	walletRepo      *repository.WalletRepository
}

func NewWithdrawalHandler(transactionRepo *repository.TransactionRepo, walletRepo *repository.WalletRepository) *WithdrawalHandler {
	return &WithdrawalHandler{
		transactionRepo: transactionRepo,
		walletRepo:      walletRepo,
	}
}
