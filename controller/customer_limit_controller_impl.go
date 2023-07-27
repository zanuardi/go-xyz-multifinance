package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/zanuardi/go-xyz-multifinance/helper"
	"github.com/zanuardi/go-xyz-multifinance/logger"
	"github.com/zanuardi/go-xyz-multifinance/model/request"
	"github.com/zanuardi/go-xyz-multifinance/model/response"
	"github.com/zanuardi/go-xyz-multifinance/service"

	"github.com/julienschmidt/httprouter"
)

type CustomerLimitControllerImpl struct {
	customerLimitService service.CustomerLimitService
}

func NewCustomerLimitController(customerLimitService service.CustomerLimitService) CustomerLimitController {
	return &CustomerLimitControllerImpl{
		customerLimitService: customerLimitService,
	}
}

func (customerLimitController *CustomerLimitControllerImpl) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	ctx := context.Background()
	logCtx := "CustomerLimitControllerImpl.Create"

	customerLimitRequest := request.CustomerLimitRequest{}
	helper.ReadFromRequestBody(r, &customerLimitRequest)

	customerLimitResponse, err := customerLimitController.customerLimitService.Create(r.Context(), customerLimitRequest)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerLimitResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerLimitController *CustomerLimitControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	ctx := context.Background()
	logCtx := "CustomerLimitControllerImpl.FindAll"

	customerLimitResponses, err := customerLimitController.customerLimitService.FindAll(r.Context())
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerLimitResponses,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerLimitController *CustomerLimitControllerImpl) FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	ctx := context.Background()
	logCtx := "CustomerLimitControllerImpl.FindById"

	customerLimitId := param.ByName("limit_id")
	id, err := strconv.Atoi(customerLimitId)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	customerLimitResponse, err := customerLimitController.customerLimitService.FindById(r.Context(), id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerLimitResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerLimitController *CustomerLimitControllerImpl) UpdateById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	ctx := context.Background()
	logCtx := "CustomerLimitControllerImpl.UpdateById"

	customerLimitRequest := request.CustomerLimitRequest{}
	helper.ReadFromRequestBody(r, &customerLimitRequest)

	customerLimitId := param.ByName("limit_id")
	id, err := strconv.Atoi(customerLimitId)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	customerLimitRequest.Id = id

	customerLimitResponse, err := customerLimitController.customerLimitService.UpdateById(r.Context(), customerLimitRequest)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerLimitResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerLimitController *CustomerLimitControllerImpl) DeleteById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	ctx := context.Background()
	logCtx := "CustomerLimitControllerImpl.DeleteById"

	customerLimitId := param.ByName("limit_id")
	id, err := strconv.Atoi(customerLimitId)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	err = customerLimitController.customerLimitService.DeleteById(r.Context(), id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)

}
