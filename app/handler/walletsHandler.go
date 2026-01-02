package handler

import (
	"context"

	"github.com/core-wallet/app/dto"
	"github.com/core-wallet/app/repository"
	"github.com/google/uuid"
)

type WalletHandler struct {
	walletRepo *repository.WalletRepository
}

func NewWalletHandler(walletRepo *repository.WalletRepository) *WalletHandler {
	return &WalletHandler{
		walletRepo: walletRepo,
	}
}

func (h *WalletHandler) CreateWallet(ctx context.Context, userId string) (string, error) {
	walletId := uuid.New().String()

	wallet, err := h.walletRepo.CreateWallet(ctx, userId, walletId)
	if err != nil {
		return "", err
	}

	return wallet.ID, nil
}

func (h *WalletHandler) GetUserBalanceByWalletId(walletId string) (resp dto.BalanceResponse, err error) {
	data, err := h.walletRepo.GetWalletById(walletId)
	if err != nil {
		return resp, err
	}
	resp.WalletId = data.ID
	resp.Balance = data.Balance
	resp.LastUpdated = data.UpdatedAt
	return resp, nil
}
