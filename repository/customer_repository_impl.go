package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/zanuardi/go-xyz-multifinance/logger"
	"github.com/zanuardi/go-xyz-multifinance/model/domain"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

func (repository *CustomerRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, customer domain.Customer) (domain.Customer, error) {
	logCtx := "customerRepository.Create"
	logger.Info(ctx, logCtx)

	query := `INSERT INTO customers(nik, full_name, legal_name,
		birth_place, birth_date, salary, ktp_photo,
		selfie_photo, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?,?)`

	now := time.Now()
	result, err := tx.ExecContext(ctx, query, customer.NIK, customer.FullName, customer.LegalName, customer.BirthPlace,
		customer.BirthDate, customer.Salary, customer.KTPPhoto, customer.SelfiePhoto, now, now)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	customer.Id = int(id)

	return customer, err

}

func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Customer, error) {
	logCtx := "customerRepository.FindAll"
	logger.Info(ctx, logCtx)

	customers := []domain.Customer{}

	query := `SELECT id, nik, full_name, legal_name, birth_place, birth_date, salary, ktp_photo,
	selfie_photo, created_at, updated_at FROM customers WHERE deleted_at IS NULL;`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer rows.Close()

	for rows.Next() {
		customer := domain.Customer{}
		fmt.Println("RES", customers)

		rows.Scan(&customer.Id, &customer.NIK, &customer.FullName, &customer.LegalName,
			&customer.BirthPlace, &customer.BirthDate, &customer.Salary, &customer.KTPPhoto,
			&customer.SelfiePhoto, &customer.CreatedAt, &customer.UpdatedAt)

		fmt.Println("RES", customers)
		customers = append(customers, customer)
		if err != nil {
			logger.Error(ctx, logCtx, err)
		}
	}
	return customers, err
}

func (repository *CustomerRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Customer, error) {
	logCtx := "customerRepository.FindById"
	logger.Info(ctx, logCtx)

	customer := domain.Customer{}
	query := `SELECT id, nik, full_name, legal_name, birth_place, birth_date, salary, ktp_photo,
	selfie_photo, created_at, updated_at FROM customers WHERE id = ?  AND deleted_at IS NULL`

	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&customer.Id, &customer.NIK, &customer.FullName, &customer.LegalName,
			&customer.BirthPlace, &customer.BirthDate, &customer.Salary, &customer.KTPPhoto,
			&customer.SelfiePhoto, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			logger.Error(ctx, logCtx, err)
		}
		return customer, nil
	} else {
		logger.Error(ctx, logCtx, err)
		return customer, errors.New("customer not found")
	}

}

func (repository *CustomerRepositoryImpl) UpdateById(ctx context.Context, tx *sql.Tx, customer domain.Customer) (domain.Customer, error) {
	logCtx := "customerRepository.UpdateById"
	logger.Info(ctx, logCtx)

	now := time.Now()

	query := `UPDATE customers
			SET nik=?, full_name=?, legal_name=?, birth_place=?, birth_date=?,
			salary=?, ktp_photo=?, selfie_photo=?, updated_at=?
			WHERE id=?;
`

	_, err := tx.ExecContext(ctx, query, &customer.NIK, &customer.FullName, &customer.LegalName,
		&customer.BirthPlace, &customer.BirthDate, &customer.Salary, &customer.KTPPhoto,
		&customer.SelfiePhoto, now, &customer.Id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	return customer, nil

}

func (repository *CustomerRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, id int) error {
	logCtx := "customerRepository.DeleteById"
	logger.Info(ctx, logCtx)

	now := time.Now()

	query := `UPDATE customers
			SET deleted_at=? WHERE id = ?`

	_, err := tx.ExecContext(ctx, query, now, id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	return nil
}
