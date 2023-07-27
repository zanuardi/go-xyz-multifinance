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

type CustomerLimitServiceImpl struct {
	CustomerLimitRepository repository.CustomerLimitRepository
	DB                      *sql.DB
	Validate                *validator.Validate
}

func NewCustomerLimitService(customerRepository repository.CustomerLimitRepository, DB *sql.DB, validate *validator.Validate) CustomerLimitService {
	return &CustomerLimitServiceImpl{
		CustomerLimitRepository: customerRepository,
		DB:                      DB,
		Validate:                validate,
	}
}

func toCustomerLimitResponse(customer domain.CustomerLimit) response.CustomerLimitResponse {
	return response.CustomerLimitResponse{
		Id:             customer.Id,
		CustomerId:     customer.CustomerId,
		Limit1:         customer.Limit1,
		Limit2:         customer.Limit2,
		Limit3:         customer.Limit3,
		Limit4:         customer.Limit4,
		RemainingLimit: customer.RemainingLimit,
		CreatedAt:      customer.CreatedAt,
		UpdatedAt:      customer.UpdatedAt,
	}
}

func toCustomerLimitsResponse(categories []domain.CustomerLimit) []response.CustomerLimitResponse {
	var customerResponses []response.CustomerLimitResponse
	for _, customer := range categories {
		customerResponses = append(customerResponses, toCustomerLimitResponse(customer))
	}
	return customerResponses
}

func (customerService *CustomerLimitServiceImpl) Create(ctx context.Context, request request.CustomerLimitRequest) (response.CustomerLimitResponse, error) {
	logCtx := "CustomerLimitServiceImpl.Create"

	err := customerService.Validate.Struct(request)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customer := domain.CustomerLimit{
		CustomerId:     request.CustomerId,
		Limit1:         request.Limit1,
		Limit2:         request.Limit2,
		Limit3:         request.Limit3,
		Limit4:         request.Limit4,
		RemainingLimit: request.RemainingLimit,
	}

	customerRequest, err := customerService.CustomerLimitRepository.Create(ctx, tx, customer)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	return toCustomerLimitResponse(customerRequest), err

}

func (customerService *CustomerLimitServiceImpl) FindAll(ctx context.Context) ([]response.CustomerLimitResponse, error) {
	logCtx := "CustomerLimitServiceImpl.FindAll"

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customers, err := customerService.CustomerLimitRepository.FindAll(ctx, tx)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	return toCustomerLimitsResponse(customers), err
}

func (customerService *CustomerLimitServiceImpl) FindById(ctx context.Context, id int) (response.CustomerLimitResponse, error) {
	logCtx := "CustomerLimitServiceImpl.FindById"

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customer, err := customerService.CustomerLimitRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return toCustomerLimitResponse(customer), err
}

func (customerService *CustomerLimitServiceImpl) UpdateById(ctx context.Context, request request.CustomerLimitRequest) (response.CustomerLimitResponse, error) {
	logCtx := "CustomerLimitServiceImpl.UpdateById"

	err := customerService.Validate.Struct(request)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customer, err := customerService.CustomerLimitRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	req := domain.CustomerLimit{
		Id:             customer.Id,
		CustomerId:     request.CustomerId,
		Limit1:         request.Limit1,
		Limit2:         request.Limit2,
		Limit3:         request.Limit3,
		Limit4:         request.Limit4,
		RemainingLimit: request.RemainingLimit,
	}
	res, err := customerService.CustomerLimitRepository.UpdateById(ctx, tx, req)

	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	return toCustomerLimitResponse(res), nil
}

func (customerService *CustomerLimitServiceImpl) DeleteById(ctx context.Context, id int) error {
	logCtx := "CustomerLimitServiceImpl.DeleteById"

	tx, err := customerService.DB.Begin()
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}
	defer helper.CommitOrRollback(tx)

	customer, err := customerService.CustomerLimitRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	customerService.CustomerLimitRepository.DeleteById(ctx, tx, customer.Id)

	return nil
}
