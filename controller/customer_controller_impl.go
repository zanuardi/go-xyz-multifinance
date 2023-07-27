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

type CustomerControllerImpl struct {
	customerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		customerService: customerService,
	}
}

func (customerController *CustomerControllerImpl) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	ctx := context.Background()
	logCtx := "CustomerControllerImpl.Create"
	logger.Info(ctx, logCtx)

	customerRequest := request.CustomerRequest{}
	helper.ReadFromRequestBody(r, &customerRequest)

	customerResponse, err := customerController.customerService.Create(r.Context(), customerRequest)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerController *CustomerControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	ctx := context.Background()
	logCtx := "CustomerControllerImpl.FindAll"

	customerResponses, err := customerController.customerService.FindAll(r.Context())
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponses,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerController *CustomerControllerImpl) FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	ctx := context.Background()
	logCtx := "CustomerControllerImpl.FindById"

	customerId := param.ByName("customer_id")
	id, err := strconv.Atoi(customerId)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	customerResponse, err := customerController.customerService.FindById(r.Context(), id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerController *CustomerControllerImpl) UpdateById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	ctx := context.Background()
	logCtx := "CustomerControllerImpl.UpdateById"

	customerRequest := request.CustomerRequest{}
	helper.ReadFromRequestBody(r, &customerRequest)

	customerId := param.ByName("customer_id")
	id, err := strconv.Atoi(customerId)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	customerRequest.Id = id

	customerResponse, err := customerController.customerService.UpdateById(r.Context(), customerRequest)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerController *CustomerControllerImpl) DeleteById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	ctx := context.Background()
	logCtx := "CustomerControllerImpl.DeleteById"

	customerId := param.ByName("customer_id")
	id, err := strconv.Atoi(customerId)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	err = customerController.customerService.DeleteById(r.Context(), id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)

}
