package handler

import (
	"context"
	"errors"

	"github.com/core-wallet/app/dto"
	"github.com/core-wallet/app/models"
	"github.com/core-wallet/app/repository"
	"github.com/core-wallet/app/utils"
)

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

func (h *WithdrawalHandler) RequestWithdrawal(ctx context.Context, request dto.WithdrawalRequest) (string, error) {
	wallet, err := h.walletRepo.GetWalletById(request.WalletID)
	if err != nil {
		utils.ErrorLog("failed to get wallet", err, true)
		return "", err
	}

	if wallet.Balance < request.Amount || request.Amount <= 0 {
		return "", errors.New("insufficient balance")
	}

	data := models.Transaction{
		WalletID:      request.WalletID,
		Type:          models.TransactionTypeWithdraw,
		Amount:        request.Amount,
		PaymentMethod: request.PaymentMethod,
		BalanceBefore: wallet.Balance,
		BalanceAfter:  wallet.Balance - request.Amount,
		Status:        models.TransactionStatusPending,
	}

	if err = h.transactionRepo.CreateTransaction(ctx, &data); err != nil {
		return "", err
	}
	return data.ID, nil
}

func (h *WithdrawalHandler) ConfirmWithdrawal(ctx context.Context, request dto.ConfirmWithdrawalRequest) error {
	return h.walletRepo.ConfirmTransaction(ctx, request.TransactionID, request.WalletID, h.transactionRepo)
}

func (h *WithdrawalHandler) CheckStatusTransaction(ctx context.Context, request dto.CheckStatusTopUpRequest) (*models.Transaction, error) {
	return h.transactionRepo.GetTransactionById(ctx, request.TransactionID, false)
}
