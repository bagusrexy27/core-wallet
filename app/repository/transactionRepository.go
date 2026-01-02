package repository

import (
	"context"
	"errors"

	"github.com/core-wallet/app/models"
	"github.com/core-wallet/app/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepo {
	return &TransactionRepo{db: db}
}

func (r *TransactionRepo) GetTransactionById(ctx context.Context, tx *gorm.DB, transactionId string, lock bool) (*models.Transaction, error) {
	if tx == nil {
		tx = r.db
	}

	query := tx.WithContext(ctx)

	if lock {
		query = query.Clauses(clause.Locking{Strength: "UPDATE"})
	}

	var transaction models.Transaction
	if err := query.Where("id = ?", transactionId).First(&transaction).Error; err != nil {
		utils.ErrorLog("failed to get transaction by id", err, true)
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepo) CreateTransaction(ctx context.Context, tx *gorm.DB, transaction *models.Transaction) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(transaction).Error; err != nil {
		utils.ErrorLog("failed to create transaction", err, true)
		return err
	}
	return nil
}

func (r *TransactionRepo) UpdateTransaction(ctx context.Context, tx *gorm.DB, transactionId string, status models.TransactionStatus) error {
	if tx == nil {
		tx = r.db
	}

	var transaction models.Transaction
	if err := tx.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", transactionId).
		First(&transaction).Error; err != nil {
		utils.ErrorLog("failed to get transaction for update", err, true)
		return err
	}

	// Update status
	transaction.Status = status

	if err := tx.Save(&transaction).Error; err != nil {
		utils.ErrorLog("failed to update transaction", err, true)
		return err
	}
	return nil
}

func (r *TransactionRepo) RejectTransaction(ctx context.Context, transactionId string) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var transaction models.Transaction
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", transactionId).
		First(&transaction).Error; err != nil {
		tx.Rollback()
		utils.ErrorLog("failed to get transaction", err, true)
		return err
	}

	if transaction.Status != models.TransactionStatusPending {
		tx.Rollback()
		return errors.New("can only reject pending transactions")
	}

	transaction.Status = models.TransactionStatusFailed

	if err := tx.Save(&transaction).Error; err != nil {
		tx.Rollback()
		utils.ErrorLog("failed to reject transaction", err, true)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
