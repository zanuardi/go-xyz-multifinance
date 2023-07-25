package repository

import (
	"context"
	"database/sql"

	"github.com/zanuardi/go-xyz-multifinance/model/domain"
)

type CustomerLimitRepository interface {
	Create(ctx context.Context, tx *sql.Tx, customerLimit domain.CustomerLimit) (domain.CustomerLimit, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.CustomerLimit, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.CustomerLimit, error)
	UpdateById(ctx context.Context, tx *sql.Tx, customerLimit domain.CustomerLimit) (domain.CustomerLimit, error)
	DeleteById(ctx context.Context, tx *sql.Tx, id int) error
}
