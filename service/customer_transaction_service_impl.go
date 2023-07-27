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

type CustomerTransactionServiceImpl struct {
	CustomerTransactionRepository repository.CustomerTransactionRepository
	CustomerLimitRepository       repository.CustomerLimitRepository
	CustomerInstallmentRepository repository.CustomerInstallmentRepository
	DB                            *sql.DB
	Validate                      *validator.Validate
}

func NewCustomerTransactionService(customerRepository repository.CustomerTransactionRepository, customerLimitRepository repository.CustomerLimitRepository, customerInstallmentRepository repository.CustomerInstallmentRepository, DB *sql.DB, validate *validator.Validate) CustomerTransactionService {
	return &CustomerTransactionServiceImpl{
		CustomerTransactionRepository: customerRepository,
		CustomerLimitRepository:       customerLimitRepository,
		CustomerInstallmentRepository: customerInstallmentRepository,
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

	customer := domain.CustomerTransaction{
		CustomerId:        request.CustomerId,
		ContractNumber:    request.ContractNumber,
		OTRPrice:          request.OTRPrice,
		AdminFee:          request.AdminFee,
		InstallmentAmount: request.InstallmentAmount,
		InterestAmount:    request.InterestAmount,
		AssetName:         request.AssetName,
		Status:            "PENDING",
	}

	totalAmount := request.OTRPrice + request.AdminFee

	resultChan := make(chan bool)
	go customerService.checkCustomerLimit(ctx, request.CustomerId, totalAmount, request.Tenor, resultChan)

	valid := <-resultChan

	if valid {
		customerRequest, err := customerService.CustomerTransactionRepository.Create(ctx, tx, customer)
		if err != nil {
			logger.Error(ctx, logCtx, err)
		}
		return toCustomerTransactionResponse(customerRequest), err
	} else {
		logger.Error(ctx, logCtx, err)
		return response.CustomerTransactionResponse{}, err
	}

}

func (customerService *CustomerTransactionServiceImpl) checkCustomerLimit(ctx context.Context, customerID int, amount float32, tenor int, resultChan chan bool) {
	logCtx := "CustomerTransactionServiceImpl.checkCustomerLimit"
	logger.Info(ctx, logCtx)

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	checkLimit, err := customerService.CustomerLimitRepository.FindByCustomerId(ctx, tx, customerID)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	if err != nil {
		logger.Error(ctx, logCtx+"Error retrieving limit for customer", err)
		resultChan <- false
		return
	}

	var limit float32
	if tenor == 1 {
		limit = checkLimit.Limit1
	} else if tenor == 2 {
		limit = checkLimit.Limit2
	} else if tenor == 3 {
		limit = checkLimit.Limit3
	} else {
		limit = checkLimit.Limit4
	}

	if limit >= amount {
		resultChan <- true
	} else {
		resultChan <- false
	}
}

func (customerService *CustomerTransactionServiceImpl) FindAll(ctx context.Context) ([]response.CustomerTransactionResponse, error) {
	logCtx := "CustomerTransactionServiceImpl.FindAll"

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customers, err := customerService.CustomerTransactionRepository.FindAll(ctx, tx)
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

func (customerService *CustomerTransactionServiceImpl) FindByCustomerId(ctx context.Context, customerId int) (response.CustomerTransactionResponse, error) {
	logCtx := "CustomerTransactionServiceImpl.FindById"

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customer, err := customerService.CustomerTransactionRepository.FindByCustomerId(ctx, tx, customerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return toCustomerTransactionResponse(customer), err
}
