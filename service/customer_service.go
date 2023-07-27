package service

import (
	"context"

	"github.com/zanuardi/go-xyz-multifinance/model/request"
	"github.com/zanuardi/go-xyz-multifinance/model/response"
)

type CustomerService interface {
	Create(ctx context.Context, request request.CustomerRequest) (response.CustomerResponse, error)
	FindAll(ctx context.Context) ([]response.CustomerResponse, error)
	FindById(ctx context.Context, id int) (response.CustomerResponse, error)
	UpdateById(ctx context.Context, request request.CustomerRequest) (response.CustomerResponse, error)
	DeleteById(ctx context.Context, id int) error
}
