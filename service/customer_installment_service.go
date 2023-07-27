package service

import (
	"context"

	"github.com/zanuardi/go-xyz-multifinance/model/request"
	"github.com/zanuardi/go-xyz-multifinance/model/response"
)

type CustomerInstallmentService interface {
	Create(ctx context.Context, request request.CustomerInstallmentRequest) (response.CustomerInstallmentResponse, error)
	FindAll(ctx context.Context) ([]response.CustomerInstallmentResponse, error)
	FindById(ctx context.Context, id int) (response.CustomerInstallmentResponse, error)
	UpdateById(ctx context.Context, request request.CustomerInstallmentRequest) (response.CustomerInstallmentResponse, error)
	DeleteById(ctx context.Context, id int) error
}
