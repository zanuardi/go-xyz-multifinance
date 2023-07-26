package service

import (
	"context"

	"github.com/zanuardi/go-xyz-multifinance/model/request"
	"github.com/zanuardi/go-xyz-multifinance/model/response"
)

type CustomerLimitService interface {
	Create(ctx context.Context, request request.CustomerLimitRequest) (response.CustomerLimitResponse, error)
	FindAll(ctx context.Context) ([]response.CustomerLimitResponse, error)
	FindById(ctx context.Context, id int) (response.CustomerLimitResponse, error)
	UpdateById(ctx context.Context, request request.CustomerLimitRequest) (response.CustomerLimitResponse, error)
	DeleteById(ctx context.Context, id int) error
}
