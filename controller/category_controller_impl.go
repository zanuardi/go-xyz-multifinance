package controller

import (
	"net/http"
	"strconv"

	"github.com/zanuardi/go-xyz-multifinance/helper"
	"github.com/zanuardi/go-xyz-multifinance/model/request"
	"github.com/zanuardi/go-xyz-multifinance/model/response"
	"github.com/zanuardi/go-xyz-multifinance/service"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		categoryService: categoryService,
	}
}

func (categoryController *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	categoryRequest := request.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryRequest)

	categoryResponse, err := categoryController.categoryService.Create(r.Context(), categoryRequest)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (categoryController *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	categoryResponses, err := categoryController.categoryService.FindAll(r.Context())
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (categoryController *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	categoryId := param.ByName("category_id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse, err := categoryController.categoryService.FindById(r.Context(), id)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (categoryController *CategoryControllerImpl) UpdateById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	categoryRequest := request.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryRequest)

	categoryId := param.ByName("category_id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryRequest.Id = id

	categoryResponse, err := categoryController.categoryService.UpdateById(r.Context(), categoryRequest)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (categoryController *CategoryControllerImpl) DeleteById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	categoryId := param.ByName("category_id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	err = categoryController.categoryService.DeleteById(r.Context(), id)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)

}
