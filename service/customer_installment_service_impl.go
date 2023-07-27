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

type CustomerInstallmentServiceImpl struct {
	CustomerInstallmentRepository repository.CustomerInstallmentRepository
	DB                            *sql.DB
	Validate                      *validator.Validate
}

func NewCustomerInstallmentService(customerRepository repository.CustomerInstallmentRepository, DB *sql.DB, validate *validator.Validate) CustomerInstallmentService {
	return &CustomerInstallmentServiceImpl{
		CustomerInstallmentRepository: customerRepository,
		DB:                            DB,
		Validate:                      validate,
	}
}

func toCustomerInstallmentResponse(customer domain.CustomerInstallment) response.CustomerInstallmentResponse {
	return response.CustomerInstallmentResponse{
		Id:                    customer.Id,
		CustomerTransactionId: customer.CustomerTransactionId,
		CustomerLimitId:       customer.CustomerLimitId,
		Tenor:                 customer.Tenor,
		TotalAmounts:          customer.TotalAmounts,
		RemainingAmounts:      customer.RemainingAmounts,
		RemainingLimit:        customer.RemainingLimit,
		CreatedAt:             customer.CreatedAt,
		UpdatedAt:             customer.UpdatedAt,
	}
}

func toCustomerInstallmentsResponse(categories []domain.CustomerInstallment) []response.CustomerInstallmentResponse {
	var customerResponses []response.CustomerInstallmentResponse
	for _, customer := range categories {
		customerResponses = append(customerResponses, toCustomerInstallmentResponse(customer))
	}
	return customerResponses
}

func (customerService *CustomerInstallmentServiceImpl) Create(ctx context.Context, request request.CustomerInstallmentRequest) (response.CustomerInstallmentResponse, error) {
	logCtx := "CustomerInstallmentServiceImpl.Create"

	err := customerService.Validate.Struct(request)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customer := domain.CustomerInstallment{
		CustomerTransactionId: request.CustomerTransactionId,
		CustomerLimitId:       request.CustomerLimitId,
		Tenor:                 request.Tenor,
		TotalAmounts:          request.TotalAmounts,
		RemainingAmounts:      request.RemainingAmounts,
		RemainingLimit:        request.RemainingLimit,
	}

	customerRequest, err := customerService.CustomerInstallmentRepository.Create(ctx, tx, customer)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	return toCustomerInstallmentResponse(customerRequest), err

}

func (customerService *CustomerInstallmentServiceImpl) FindAll(ctx context.Context) ([]response.CustomerInstallmentResponse, error) {
	logCtx := "CustomerInstallmentServiceImpl.FindAll"

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customers, err := customerService.CustomerInstallmentRepository.FindAll(ctx, tx)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	return toCustomerInstallmentsResponse(customers), err
}

func (customerService *CustomerInstallmentServiceImpl) FindById(ctx context.Context, id int) (response.CustomerInstallmentResponse, error) {
	logCtx := "CustomerInstallmentServiceImpl.FindById"

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customer, err := customerService.CustomerInstallmentRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return toCustomerInstallmentResponse(customer), err
}

func (customerService *CustomerInstallmentServiceImpl) UpdateById(ctx context.Context, request request.CustomerInstallmentRequest) (response.CustomerInstallmentResponse, error) {
	logCtx := "CustomerInstallmentServiceImpl.UpdateById"

	err := customerService.Validate.Struct(request)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customer, err := customerService.CustomerInstallmentRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	req := domain.CustomerInstallment{
		Id:                    customer.Id,
		CustomerTransactionId: request.CustomerTransactionId,
		CustomerLimitId:       request.CustomerLimitId,
		Tenor:                 request.Tenor,
		TotalAmounts:          request.TotalAmounts,
		RemainingAmounts:      request.RemainingAmounts,
		RemainingLimit:        request.RemainingLimit,
	}
	res, err := customerService.CustomerInstallmentRepository.UpdateById(ctx, tx, req)

	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	return toCustomerInstallmentResponse(res), nil
}

func (customerService *CustomerInstallmentServiceImpl) DeleteById(ctx context.Context, id int) error {
	logCtx := "CustomerInstallmentServiceImpl.DeleteById"

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customer, err := customerService.CustomerInstallmentRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	customerService.CustomerInstallmentRepository.DeleteById(ctx, tx, customer.Id)

	return nil
}
