package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/zanuardi/go-xyz-multifinance/logger"
	"github.com/zanuardi/go-xyz-multifinance/model/domain"
)

type CustomerTransactionRepositoryImpl struct {
}

func NewCustomerTransactionRepository() CustomerTransactionRepository {
	return &CustomerTransactionRepositoryImpl{}
}

func (repository *CustomerTransactionRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, customerTransaction domain.CustomerTransaction) (domain.CustomerTransaction, error) {
	logCtx := "customerTransactionRepository.Create"
	logger.Info(ctx, logCtx)

	query := `INSERT INTO customer_transactions
		(customer_id, contract_number, otr_price, admin_fee, installment_amount,
		interest_amount, asset_name, status, created_at, updated_at
		VALUES(?,?,?,?,?,?,?,?,?,?);`

	now := time.Now()
	result, err := tx.ExecContext(ctx, query, customerTransaction.CustomerId, customerTransaction.ContractNumber, customerTransaction.OTRPrice, customerTransaction.AdminFee,
		customerTransaction.InstallmentAmount, customerTransaction.InterestAmount, customerTransaction.AssetName, customerTransaction.Status, now, now)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	customerTransaction.Id = int(id)

	return customerTransaction, err

}

func (repository *CustomerTransactionRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.CustomerTransaction, error) {
	logCtx := "customerTransactionRepository.FindAll"
	logger.Info(ctx, logCtx)

	customers := []domain.CustomerTransaction{}

	query := `SELECT id, customer_id, contract_number, otr_price, admin_fee, installment_amount,
		interest_amount, asset_name, status, created_at, updated_at, deleted_at
		FROM customer_transactions WHERE deleted_at IS NULL;`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer rows.Close()

	for rows.Next() {
		customerTransaction := domain.CustomerTransaction{}
		rows.Scan(&customerTransaction.Id, &customerTransaction.CustomerId, &customerTransaction.ContractNumber, &customerTransaction.OTRPrice,
			&customerTransaction.AdminFee, &customerTransaction.InstallmentAmount, &customerTransaction.InterestAmount, &customerTransaction.AssetName,
			&customerTransaction.Status, &customerTransaction.CreatedAt, &customerTransaction.UpdatedAt)

		customers = append(customers, customerTransaction)
		if err != nil {
			logger.Error(ctx, logCtx, err)
		}
	}
	return customers, err
}

func (repository *CustomerTransactionRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.CustomerTransaction, error) {
	logCtx := "customerTransactionRepository.FindById"
	logger.Info(ctx, logCtx)

	customerTransaction := domain.CustomerTransaction{}
	query := `SELECT id, customer_id, contract_number, otr_price, admin_fee,
		installment_amount,	interest_amount, asset_name, status, created_at, updated_at
		FROM customer_transactions WHERE id = ?  AND deleted_at IS NULL`

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&customerTransaction.Id, &customerTransaction.CustomerId, &customerTransaction.ContractNumber, &customerTransaction.OTRPrice,
			&customerTransaction.AdminFee, &customerTransaction.InstallmentAmount, &customerTransaction.InterestAmount, &customerTransaction.AssetName,
			&customerTransaction.Status, &customerTransaction.CreatedAt, &customerTransaction.UpdatedAt)
		if err != nil {
			logger.Error(ctx, logCtx, err)
		}
		return customerTransaction, nil
	} else {
		logger.Error(ctx, logCtx, err)
		return customerTransaction, errors.New("customerTransaction not found")
	}

}

func (repository *CustomerTransactionRepositoryImpl) UpdateById(ctx context.Context, tx *sql.Tx, customerTransaction domain.CustomerTransaction) (domain.CustomerTransaction, error) {
	logCtx := "customerTransactionRepository.UpdateById"
	logger.Info(ctx, logCtx)

	query := `UPDATE customer_transactions
			SET customer_id=?, contract_number=?, otr_price=?, admin_fee=?,
			installment_amount=?, interest_amount=?, asset_name=?, status=?,
			updated_at=? WHERE id=?;`

	now := time.Now()
	_, err := tx.ExecContext(ctx, query, &customerTransaction.CustomerId, &customerTransaction.ContractNumber, &customerTransaction.OTRPrice,
		&customerTransaction.AdminFee, &customerTransaction.InstallmentAmount, &customerTransaction.InterestAmount, &customerTransaction.AssetName,
		&customerTransaction.Status, now, &customerTransaction.Id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	return customerTransaction, nil

}

func (repository *CustomerTransactionRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, id int) error {
	logCtx := "customerTransactionRepository.DeleteById"
	logger.Info(ctx, logCtx)

	query := `UPDATE customer_transactions
			SET deleted_at=? WHERE id = ?`

	now := time.Now()

	_, err := tx.ExecContext(ctx, query, now, id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	return nil
}
