package repository

import (
	"context"
	"database/sql"

	"github.com/zanuardi/go-xyz-multifinance/model/domain"
)

type CustomerTransactionRepository interface {
	Create(ctx context.Context, tx *sql.Tx, customerTransaction domain.CustomerTransaction) (domain.CustomerTransaction, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.CustomerTransaction, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.CustomerTransaction, error)
	FindByCustomerId(ctx context.Context, tx *sql.Tx, customerId int) (domain.CustomerTransaction, error)
}
