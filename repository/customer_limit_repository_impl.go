package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/zanuardi/go-xyz-multifinance/logger"
	"github.com/zanuardi/go-xyz-multifinance/model/domain"
)

type CustomerLimitRepositoryImpl struct {
}

func NewCustomerLimitRepository() CustomerLimitRepository {
	return &CustomerLimitRepositoryImpl{}
}

func (repository *CustomerLimitRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, customerLimit domain.CustomerLimit) (domain.CustomerLimit, error) {
	logCtx := "customerLimitRepository.Create"
	logger.Info(ctx, logCtx)

	query := `INSERT INTO customer_limits
			(customer_id, limit_1, limit_2, limit_3, limit_4, remaining_limit,
			created_at, updated_at
 			VALUES (?,?,?,?,?,?,?,?)`

	now := time.Now()
	result, err := tx.ExecContext(ctx, query, customerLimit.CustomerId, customerLimit.Limit1, customerLimit.Limit2, customerLimit.Limit3,
		customerLimit.Limit4, customerLimit.RemainingLimit, now, now)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	customerLimit.Id = int(id)

	return customerLimit, err

}

func (repository *CustomerLimitRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.CustomerLimit, error) {
	logCtx := "customerLimitRepository.FindAll"
	logger.Info(ctx, logCtx)

	customerLimits := []domain.CustomerLimit{}

	query := `SELECT id, customer_id, limit_1, limit_2, limit_3, limit_4,
			remaining_limit, created_at, updated_at, deleted_at
			FROM customer_limits;`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer rows.Close()

	for rows.Next() {
		customerLimit := domain.CustomerLimit{}
		rows.Scan(&customerLimit.Id, &customerLimit.CustomerId, &customerLimit.Limit1, &customerLimit.Limit2,
			&customerLimit.Limit3, &customerLimit.Limit4, &customerLimit.RemainingLimit,
			&customerLimit.CreatedAt, &customerLimit.UpdatedAt)

		customerLimits = append(customerLimits, customerLimit)
		if err != nil {
			logger.Error(ctx, logCtx, err)
		}
	}
	return customerLimits, err
}

func (repository *CustomerLimitRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.CustomerLimit, error) {
	logCtx := "customerLimitRepository.FindById"
	logger.Info(ctx, logCtx)

	customerLimit := domain.CustomerLimit{}
	query := `SELECT id, customer_id, limit_1, limit_2, limit_3, limit_4,
			remaining_limit, created_at, updated_at, deleted_at
			FROM customer_limits WHERE id = ?`

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&customerLimit.Id, &customerLimit.CustomerId, &customerLimit.Limit1, &customerLimit.Limit2,
			&customerLimit.Limit3, &customerLimit.Limit4, &customerLimit.RemainingLimit,
			&customerLimit.CreatedAt, &customerLimit.UpdatedAt)
		if err != nil {
			logger.Error(ctx, logCtx, err)
		}
		return customerLimit, nil
	} else {
		logger.Error(ctx, logCtx, err)
		return customerLimit, errors.New("customerLimit not found")
	}

}

func (repository *CustomerLimitRepositoryImpl) UpdateById(ctx context.Context, tx *sql.Tx, customerLimit domain.CustomerLimit) (domain.CustomerLimit, error) {
	logCtx := "customerLimitRepository.UpdateById"
	logger.Info(ctx, logCtx)

	now := time.Now()

	query := `UPDATE customer_limits
			SET customer_id=?, limit_1=?, limit_2=?, limit_3=?, limit_4=?,
			remaining_limit=?, updated_at=?
			WHERE id=?;`

	_, err := tx.ExecContext(ctx, query, &customerLimit.CustomerId, &customerLimit.Limit1, &customerLimit.Limit2,
		&customerLimit.Limit3, &customerLimit.Limit4, &customerLimit.RemainingLimit, now, &customerLimit.Id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	return customerLimit, nil

}

func (repository *CustomerLimitRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, id int) error {
	logCtx := "customerLimitRepository.DeleteById"
	logger.Info(ctx, logCtx)

	now := time.Now()

	query := `UPDATE customer_limits
			SET deleted_at=? WHERE id = ?`

	_, err := tx.ExecContext(ctx, query, now, id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	return nil
}
