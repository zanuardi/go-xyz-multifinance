package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/zanuardi/go-xyz-multifinance/helper"
	"github.com/zanuardi/go-xyz-multifinance/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error) {
	SQL := "INSERT INTO categories(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category, err

}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error) {
	categories := []domain.Category{}

	query := "SELECT id, name FROM categories"

	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		category := domain.Category{}
		rows.Scan(&category.Id, &category.Name)
		categories = append(categories, category)
		helper.PanicIfError(err)
	}
	return categories, err
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error) {
	category := domain.Category{}
	query := "SELECT id, name FROM categories WHERE id = ?"

	rows, err := tx.QueryContext(ctx, query, id)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}

}

func (repository *CategoryRepositoryImpl) UpdateById(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error) {

	query := "UPDATE categories SET name = ? WHERE id = ?"

	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	helper.PanicIfError(err)

	return category, nil

}

func (repository *CategoryRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, id int) error {

	query := "DELETE FROM categories  WHERE id = ?"

	_, err := tx.ExecContext(ctx, query, id)
	helper.PanicIfError(err)

	return nil
}
