package repository

import (
	"context"
	"database/sql"

	"github.com/zanuardi/go-xyz-multifinance/model/domain"
)

type CustomerInstallmentRepository interface {
	Create(ctx context.Context, tx *sql.Tx, customerInstallment domain.CustomerInstallment) (domain.CustomerInstallment, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.CustomerInstallment, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.CustomerInstallment, error)
	UpdateById(ctx context.Context, tx *sql.Tx, customerInstallment domain.CustomerInstallment) (domain.CustomerInstallment, error)
	DeleteById(ctx context.Context, tx *sql.Tx, id int) error
}
