package service

import (
	"context"

	"github.com/zanuardi/go-xyz-multifinance/model/request"
	"github.com/zanuardi/go-xyz-multifinance/model/response"
)

type CustomerTransactionService interface {
	Create(ctx context.Context, request request.CustomerTransactionRequest) (response.CustomerTransactionResponse, error)
	FindAll(ctx context.Context) ([]response.CustomerTransactionResponse, error)
	FindById(ctx context.Context, id int) (response.CustomerTransactionResponse, error)
	FindByCustomerId(ctx context.Context, customerId int) (response.CustomerTransactionResponse, error)
}
