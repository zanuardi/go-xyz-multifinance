package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/zanuardi/go-xyz-multifinance/logger"
	"github.com/zanuardi/go-xyz-multifinance/model/domain"
)

type CustomerInstallmentRepositoryImpl struct {
}

func NewCustomerInstallmentRepository() CustomerInstallmentRepository {
	return &CustomerInstallmentRepositoryImpl{}
}

func (repository *CustomerInstallmentRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, customerInstallment domain.CustomerInstallment) (domain.CustomerInstallment, error) {
	logCtx := "customerInstallmentRepository.Create"
	logger.Info(ctx, logCtx)

	query := `INSERT INTO customer_installments
		(customer_transaction_id, customer_limit_id, tenor, total_amounts, remaining_amount, created_at, updated_at)
		VALUES(?,?,?,?,?,?,?);`

	now := time.Now()
	result, err := tx.ExecContext(ctx, query, customerInstallment.CustomerTransactionId, customerInstallment.CustomerTransactionId,
		customerInstallment.Tenor, customerInstallment.TotalAmounts, customerInstallment.RemainingAmounts,
		customerInstallment.RemainingLimit, now, now)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	customerInstallment.Id = int(id)

	return customerInstallment, err

}

func (repository *CustomerInstallmentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.CustomerInstallment, error) {
	logCtx := "customerInstallmentRepository.FindAll"
	logger.Info(ctx, logCtx)

	res := []domain.CustomerInstallment{}

	query := `SELECT id, customer_transaction_id, customer_limit_id, tenor,
	total_amounts, remaining_amount, created_at, updated_at, deleted_at
	FROM customer_installments;`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer rows.Close()

	for rows.Next() {
		customerInstallment := domain.CustomerInstallment{}
		rows.Scan(&customerInstallment.Id, &customerInstallment.CustomerTransactionId, &customerInstallment.CustomerLimitId, &customerInstallment.Tenor,
			&customerInstallment.TotalAmounts, &customerInstallment.RemainingAmounts, &customerInstallment.RemainingLimit,
			&customerInstallment.CreatedAt, &customerInstallment.UpdatedAt)

		res = append(res, customerInstallment)
		if err != nil {
			logger.Error(ctx, logCtx, err)
		}
	}
	return res, err
}

func (repository *CustomerInstallmentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.CustomerInstallment, error) {
	logCtx := "customerInstallmentRepository.FindById"
	logger.Info(ctx, logCtx)

	customerInstallment := domain.CustomerInstallment{}
	query := `SELECT id, customer_transaction_id, customer_limit_id, tenor,
		total_amounts, remaining_amount, created_at, updated_at, deleted_at
		FROM customer_installments WHERE id = ?`

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&customerInstallment.Id, &customerInstallment.CustomerTransactionId, &customerInstallment.CustomerLimitId, &customerInstallment.Tenor,
			&customerInstallment.TotalAmounts, &customerInstallment.RemainingAmounts, &customerInstallment.RemainingLimit,
			&customerInstallment.CreatedAt, &customerInstallment.UpdatedAt)
		if err != nil {
			logger.Error(ctx, logCtx, err)
		}
		return customerInstallment, nil
	} else {
		logger.Error(ctx, logCtx, err)
		return customerInstallment, errors.New("customerInstallment not found")
	}

}

func (repository *CustomerInstallmentRepositoryImpl) UpdateById(ctx context.Context, tx *sql.Tx, customerInstallment domain.CustomerInstallment) (domain.CustomerInstallment, error) {
	logCtx := "customerInstallmentRepository.UpdateById"
	logger.Info(ctx, logCtx)

	query := `UPDATE customer_installments
			SET customer_transaction_id=?, customer_limit_id=?, tenor=?, total_amounts=?,
			remaining_amount=?, updated_at=?
			WHERE id=?;`

	now := time.Now()
	_, err := tx.ExecContext(ctx, query, &customerInstallment.CustomerTransactionId, &customerInstallment.CustomerLimitId, &customerInstallment.Tenor,
		&customerInstallment.TotalAmounts, &customerInstallment.RemainingAmounts, &customerInstallment.RemainingLimit,
		now, &customerInstallment.Id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	return customerInstallment, nil

}

func (repository *CustomerInstallmentRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, id int) error {
	logCtx := "customerInstallmentRepository.DeleteById"
	logger.Info(ctx, logCtx)

	query := `UPDATE customer_installments
			SET deleted_at=? WHERE id = ?`

	now := time.Now()

	_, err := tx.ExecContext(ctx, query, now, id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	return nil
}
