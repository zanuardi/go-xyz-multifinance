package service

import (
	"context"
	"database/sql"

	"github.com/zanuardi/go-xyz-multifinance/exception"
	"github.com/zanuardi/go-xyz-multifinance/helper"
	"github.com/zanuardi/go-xyz-multifinance/model/domain"
	"github.com/zanuardi/go-xyz-multifinance/model/request"
	"github.com/zanuardi/go-xyz-multifinance/model/response"
	"github.com/zanuardi/go-xyz-multifinance/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (categoryService *CategoryServiceImpl) Create(ctx context.Context, request request.CategoryCreateRequest) (response.CategoryResponse, error) {
	err := categoryService.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	categoryRequest, err := categoryService.CategoryRepository.Create(ctx, tx, category)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(categoryRequest), err

}

func (categoryService *CategoryServiceImpl) FindAll(ctx context.Context) ([]response.CategoryResponse, error) {
	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories, err := categoryService.CategoryRepository.FindAll(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToCategoriesResponse(categories), err
}

func (categoryService *CategoryServiceImpl) FindById(ctx context.Context, id int) (response.CategoryResponse, error) {
	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := categoryService.CategoryRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToCategoryResponse(category), err
}

func (categoryService *CategoryServiceImpl) UpdateById(ctx context.Context, request request.CategoryUpdateRequest) (response.CategoryResponse, error) {

	err := categoryService.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	findById, err := categoryService.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	findById.Name = request.Name

	category, err := categoryService.CategoryRepository.UpdateById(ctx, tx, findById)

	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category), nil
}

func (categoryService *CategoryServiceImpl) DeleteById(ctx context.Context, id int) error {
	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := categoryService.CategoryRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	categoryService.CategoryRepository.DeleteById(ctx, tx, category.Id)

	return nil
}
