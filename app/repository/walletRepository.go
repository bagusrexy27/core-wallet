package repository

import (
	"context"
	"errors"

	"github.com/core-wallet/app/models"
	"github.com/core-wallet/app/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) CreateWallet(ctx context.Context, userId, walletId string) (*models.Wallet, error) {
	initialBalance := int64(0)

	initialChecksum := utils.CalculateChecksum(userId, initialBalance)

	wallet := &models.Wallet{
		ID:       walletId,
		UserId:   userId,
		Balance:  initialBalance,
		Checksum: initialChecksum,
	}

	if err := r.db.WithContext(ctx).Create(wallet).Error; err != nil {
		utils.ErrorLog("failed to create wallet", err, true)
		return nil, err
	}

	return wallet, nil
}

func (r *WalletRepository) GetWalletById(walletId string) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := r.db.Where("id = ?", walletId).First(&wallet).Error; err != nil {
		utils.ErrorLog("failed to get wallet by id", err, true)
		return nil, err
	}
	return &wallet, nil
}

func (r *WalletRepository) ConfirmTransaction(ctx context.Context, transactionID, walletID string, transactionRepo *TransactionRepo) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	transaction, err := transactionRepo.GetTransactionById(ctx, transactionID, true)
	if err != nil {
		tx.Rollback()
		return err
	}

	if transaction.Status != models.TransactionStatusPending {
		tx.Rollback()
		return errors.New("transaction is not pending")
	}

	if transaction.WalletID != walletID {
		tx.Rollback()
		return errors.New("wallet ID mismatch")
	}

	var wallet models.Wallet
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", walletID).
		First(&wallet).Error; err != nil {
		tx.Rollback()
		utils.ErrorLog("failed to get wallet", err, true)
		return err
	}

	if !utils.ValidateChecksum(wallet.UserId, wallet.Balance, wallet.Checksum) {
		tx.Rollback()
		err := errors.New("checksum validation failed - wallet data may be corrupted")
		utils.ErrorLog("checksum mismatch", err, true)
		return err
	}

	wallet.Balance += transaction.Amount
	wallet.Checksum = utils.CalculateChecksum(wallet.UserId, wallet.Balance)

	if err := tx.Save(&wallet).Error; err != nil {
		tx.Rollback()
		utils.ErrorLog("failed to save wallet", err, true)
		return err
	}

	if err := transactionRepo.UpdateTransaction(ctx, tx, transaction.ID, models.TransactionStatusSuccess); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		utils.ErrorLog("failed to commit transaction", err, true)
		return err
	}

	return nil
}
