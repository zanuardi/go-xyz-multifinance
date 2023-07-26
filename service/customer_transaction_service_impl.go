package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zanuardi/go-xyz-multifinance/exception"
	"github.com/zanuardi/go-xyz-multifinance/helper"
	"github.com/zanuardi/go-xyz-multifinance/logger"
	"github.com/zanuardi/go-xyz-multifinance/model/domain"
	"github.com/zanuardi/go-xyz-multifinance/model/request"
	"github.com/zanuardi/go-xyz-multifinance/model/response"
	"github.com/zanuardi/go-xyz-multifinance/repository"

	"github.com/go-playground/validator/v10"
)

type CustomerTransactionServiceImpl struct {
	CustomerTransactionRepository repository.CustomerTransactionRepository
	DB                            *sql.DB
	Validate                      *validator.Validate
}

func NewCustomerTransactionService(customerRepository repository.CustomerTransactionRepository, DB *sql.DB, validate *validator.Validate) CustomerTransactionService {
	return &CustomerTransactionServiceImpl{
		CustomerTransactionRepository: customerRepository,
		DB:                            DB,
		Validate:                      validate,
	}
}

func toCustomerTransactionResponse(customer domain.CustomerTransaction) response.CustomerTransactionResponse {
	return response.CustomerTransactionResponse{
		Id:                customer.Id,
		CustomerId:        customer.CustomerId,
		ContractNumber:    customer.ContractNumber,
		OTRPrice:          customer.OTRPrice,
		AdminFee:          customer.OTRPrice,
		InstallmentAmount: customer.InstallmentAmount,
		InterestAmount:    customer.InterestAmount,
		AssetName:         customer.AssetName,
		Status:            customer.Status,
		CreatedAt:         customer.CreatedAt,
		UpdatedAt:         customer.UpdatedAt,
	}
}

func toCustomerTransactionsResponse(categories []domain.CustomerTransaction) []response.CustomerTransactionResponse {
	var customerResponses []response.CustomerTransactionResponse
	for _, customer := range categories {
		customerResponses = append(customerResponses, toCustomerTransactionResponse(customer))
	}
	return customerResponses
}

func (customerService *CustomerTransactionServiceImpl) Create(ctx context.Context, request request.CustomerTransactionRequest) (response.CustomerTransactionResponse, error) {
	logCtx := "CustomerTransactionServiceImpl.Create"

	err := customerService.Validate.Struct(request)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customer := domain.CustomerTransaction{
		CustomerId:        request.CustomerId,
		ContractNumber:    request.ContractNumber,
		OTRPrice:          request.OTRPrice,
		AdminFee:          request.OTRPrice,
		InstallmentAmount: request.InstallmentAmount,
		InterestAmount:    request.InterestAmount,
		AssetName:         request.AssetName,
		Status:            request.Status,
	}

	customerRequest, err := customerService.CustomerTransactionRepository.Create(ctx, tx, customer)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	return toCustomerTransactionResponse(customerRequest), err

}

func (customerService *CustomerTransactionServiceImpl) FindAll(ctx context.Context) ([]response.CustomerTransactionResponse, error) {
	logCtx := "CustomerTransactionServiceImpl.FindAll"
	fmt.Println("PASSED service")

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)
	fmt.Println("PASSED service")

	customers, err := customerService.CustomerTransactionRepository.FindAll(ctx, tx)
	fmt.Println(customers)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	return toCustomerTransactionsResponse(customers), err
}

func (customerService *CustomerTransactionServiceImpl) FindById(ctx context.Context, id int) (response.CustomerTransactionResponse, error) {
	logCtx := "CustomerTransactionServiceImpl.FindById"

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customer, err := customerService.CustomerTransactionRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return toCustomerTransactionResponse(customer), err
}

func (customerService *CustomerTransactionServiceImpl) UpdateById(ctx context.Context, request request.CustomerTransactionRequest) (response.CustomerTransactionResponse, error) {
	logCtx := "CustomerTransactionServiceImpl.UpdateById"

	err := customerService.Validate.Struct(request)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customer, err := customerService.CustomerTransactionRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	req := domain.CustomerTransaction{
		Id:                customer.Id,
		CustomerId:        request.CustomerId,
		ContractNumber:    request.ContractNumber,
		OTRPrice:          request.OTRPrice,
		AdminFee:          request.OTRPrice,
		InstallmentAmount: request.InstallmentAmount,
		InterestAmount:    request.InterestAmount,
		AssetName:         request.AssetName,
		Status:            request.Status,
	}
	res, err := customerService.CustomerTransactionRepository.UpdateById(ctx, tx, req)

	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	return toCustomerTransactionResponse(res), nil
}

func (customerService *CustomerTransactionServiceImpl) DeleteById(ctx context.Context, id int) error {
	logCtx := "CustomerTransactionServiceImpl.DelleteById"

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customer, err := customerService.CustomerTransactionRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	customerService.CustomerTransactionRepository.DeleteById(ctx, tx, customer.Id)

	return nil
}
