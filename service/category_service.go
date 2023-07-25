package service

import (
	"context"

	"github.com/zanuardi/go-xyz-multifinance/model/request"
	"github.com/zanuardi/go-xyz-multifinance/model/response"
)

type CategoryService interface {
	Create(ctx context.Context, request request.CategoryCreateRequest) (response.CategoryResponse, error)
	FindAll(ctx context.Context) ([]response.CategoryResponse, error)
	FindById(ctx context.Context, id int) (response.CategoryResponse, error)
	UpdateById(ctx context.Context, request request.CategoryUpdateRequest) (response.CategoryResponse, error)
	DeleteById(ctx context.Context, id int) error
}
