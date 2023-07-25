package repository

import (
	"context"
	"database/sql"

	"github.com/zanuardi/go-xyz-multifinance/model/domain"
)

type CustomerRepository interface {
	Create(ctx context.Context, tx *sql.Tx, customer domain.Customer) (domain.Customer, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Customer, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Customer, error)
	UpdateById(ctx context.Context, tx *sql.Tx, customer domain.Customer) (domain.Customer, error)
	DeleteById(ctx context.Context, tx *sql.Tx, id int) error
}
