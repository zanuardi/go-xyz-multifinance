package repository

import (
	"context"
	"database/sql"

	"github.com/zanuardi/go-xyz-multifinance/model/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error)
	UpdateById(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	DeleteById(ctx context.Context, tx *sql.Tx, id int) error
}
