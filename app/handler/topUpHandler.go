package handler

import (
	"context"

	"github.com/core-wallet/app/dto"
	"github.com/core-wallet/app/models"
	"github.com/core-wallet/app/repository"
	"github.com/core-wallet/app/utils"
)

type TopUpHandler struct {
	walletRepo      *repository.WalletRepository
	transactionRepo *repository.TransactionRepo
}

func NewTopUpHandler(
	walletRepo *repository.WalletRepository,
	transactionRepo *repository.TransactionRepo,
) *TopUpHandler {
	return &TopUpHandler{
		walletRepo:      walletRepo,
		transactionRepo: transactionRepo,
	}
}

func (h *TopUpHandler) TopUpRequest(ctx context.Context, request dto.TopUpRequest) (string, error) {
	wallet, err := h.walletRepo.GetWalletById(request.WalletID)
	if err != nil {
		utils.ErrorLog("failed to get wallet", err, true)
		return "", err
	}

	transaction := models.Transaction{
		WalletID:      request.WalletID,
		Type:          models.TransactionTypeTopup,
		Amount:        request.Amount,
		BalanceBefore: wallet.Balance,
		BalanceAfter:  wallet.Balance + request.Amount,
		PaymentMethod: request.PaymentMethod,
		Status:        models.TransactionStatusPending,
	}

	if err = h.transactionRepo.CreateTransaction(ctx, nil, &transaction); err != nil {
		return "", err
	}

	return transaction.ID, nil
}

func (h *TopUpHandler) ConfirmTopUp(ctx context.Context, request dto.ConfirmTopUpRequest) error {
	return h.walletRepo.ConfirmTopUpTransaction(ctx, request.TransactionID, request.WalletID, h.transactionRepo)
}

func (h *TopUpHandler) RejectTopUp(ctx context.Context, request dto.CheckStatusTopUpRequest) error {
	return h.transactionRepo.RejectTransaction(ctx, request.TransactionID)
}

func (h *TopUpHandler) CheckStatusTopUp(ctx context.Context, request dto.CheckStatusTopUpRequest) (*models.Transaction, error) {
	return h.transactionRepo.GetTransactionById(ctx, nil, request.TransactionID, false)
}
