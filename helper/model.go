package helper

import (
	"github.com/zanuardi/go-xyz-multifinance/model/domain"
	"github.com/zanuardi/go-xyz-multifinance/model/response"
)

func ToCategoryResponse(category domain.Category) response.CategoryResponse {
	return response.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoriesResponse(categories []domain.Category) []response.CategoryResponse {
	var categoryResponses []response.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
