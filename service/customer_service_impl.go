package service

import (
	"context"
	"database/sql"

	"github.com/zanuardi/go-xyz-multifinance/exception"
	"github.com/zanuardi/go-xyz-multifinance/helper"
	"github.com/zanuardi/go-xyz-multifinance/logger"
	"github.com/zanuardi/go-xyz-multifinance/model/domain"
	"github.com/zanuardi/go-xyz-multifinance/model/request"
	"github.com/zanuardi/go-xyz-multifinance/model/response"
	"github.com/zanuardi/go-xyz-multifinance/repository"

	"github.com/go-playground/validator/v10"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCustomerService(customerRepository repository.CustomerRepository, DB *sql.DB, validate *validator.Validate) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: customerRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func toCustomerResponse(customer domain.Customer) response.CustomerResponse {
	return response.CustomerResponse{
		Id:          customer.Id,
		NIK:         customer.NIK,
		FullName:    customer.FullName,
		LegalName:   customer.LegalName,
		BirthPlace:  customer.BirthPlace,
		BirthDate:   customer.BirthDate,
		Salary:      customer.Salary,
		KTPPhoto:    customer.KTPPhoto,
		SelfiePhoto: customer.SelfiePhoto,
		CreatedAt:   customer.CreatedAt,
		UpdatedAt:   customer.UpdatedAt,
	}
}

func toCustomersResponse(categories []domain.Customer) []response.CustomerResponse {
	var customerResponses []response.CustomerResponse
	for _, customer := range categories {
		customerResponses = append(customerResponses, toCustomerResponse(customer))
	}
	return customerResponses
}

func (customerService *CustomerServiceImpl) Create(ctx context.Context, request request.CustomerRequest) (response.CustomerResponse, error) {
	logCtx := "CustomerServiceImpl.Create"
	logger.Info(ctx, logCtx)

	err := customerService.Validate.Struct(request)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customer := domain.Customer{
		NIK:         request.NIK,
		FullName:    request.FullName,
		LegalName:   request.LegalName,
		BirthPlace:  request.BirthPlace,
		BirthDate:   request.BirthDate,
		Salary:      request.Salary,
		KTPPhoto:    request.KTPPhoto,
		SelfiePhoto: request.SelfiePhoto,
	}

	customerRequest, err := customerService.CustomerRepository.Create(ctx, tx, customer)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	return toCustomerResponse(customerRequest), err

}

func (customerService *CustomerServiceImpl) FindAll(ctx context.Context) ([]response.CustomerResponse, error) {
	logCtx := "CustomerServiceImpl.FindAll"

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customers, err := customerService.CustomerRepository.FindAll(ctx, tx)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	return toCustomersResponse(customers), err
}

func (customerService *CustomerServiceImpl) FindById(ctx context.Context, id int) (response.CustomerResponse, error) {
	logCtx := "CustomerServiceImpl.FindById"

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customer, err := customerService.CustomerRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return toCustomerResponse(customer), err
}

func (customerService *CustomerServiceImpl) UpdateById(ctx context.Context, request request.CustomerRequest) (response.CustomerResponse, error) {
	logCtx := "CustomerServiceImpl.UpdateById"

	err := customerService.Validate.Struct(request)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customer, err := customerService.CustomerRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	req := domain.Customer{
		Id:          customer.Id,
		NIK:         request.NIK,
		FullName:    request.FullName,
		LegalName:   request.LegalName,
		BirthPlace:  request.BirthPlace,
		BirthDate:   request.BirthDate,
		Salary:      request.Salary,
		KTPPhoto:    request.KTPPhoto,
		SelfiePhoto: request.SelfiePhoto,
	}
	res, err := customerService.CustomerRepository.UpdateById(ctx, tx, req)

	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	return toCustomerResponse(res), nil
}

func (customerService *CustomerServiceImpl) DeleteById(ctx context.Context, id int) error {
	logCtx := "CustomerServiceImpl.DeleteById"

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customer, err := customerService.CustomerRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	customerService.CustomerRepository.DeleteById(ctx, tx, customer.Id)

	return nil
}
