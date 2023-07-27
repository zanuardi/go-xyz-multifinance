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

type CustomerTransactionControllerImpl struct {
	customerTransactionService service.CustomerTransactionService
}

func NewCustomerTransactionController(customerTransactionService service.CustomerTransactionService) CustomerTransactionController {
	return &CustomerTransactionControllerImpl{
		customerTransactionService: customerTransactionService,
	}
}

func (customerTransactionController *CustomerTransactionControllerImpl) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	ctx := context.Background()
	logCtx := "CustomerTransactionControllerImpl.Create"

	customerTransactionRequest := request.CustomerTransactionRequest{}
	helper.ReadFromRequestBody(r, &customerTransactionRequest)

	customerTransactionResponse, err := customerTransactionController.customerTransactionService.Create(r.Context(), customerTransactionRequest)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	if customerTransactionResponse.Id > 0 {
		webResponse := response.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   customerTransactionResponse,
		}
		helper.WriteToResponseBody(w, webResponse)
	} else {
		webResponse := response.WebResponse{
			Code:   400,
			Status: "Bad Request.",
			Data:   "Not enough limit.",
		}
		helper.WriteToResponseBody(w, webResponse)
	}

}

func (customerTransactionController *CustomerTransactionControllerImpl) FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	ctx := context.Background()
	logCtx := "CustomerTransactionControllerImpl.FindById"

	customerTransactionId := param.ByName("transaction_id")
	id, err := strconv.Atoi(customerTransactionId)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	customerTransactionResponse, err := customerTransactionController.customerTransactionService.FindById(r.Context(), id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerTransactionResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerTransactionController *CustomerTransactionControllerImpl) FindByCustomerId(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	ctx := context.Background()
	logCtx := "CustomerTransactionControllerImpl.FindById"

	customerTransactionId := param.ByName("customer_id")
	id, err := strconv.Atoi(customerTransactionId)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	customerTransactionResponse, err := customerTransactionController.customerTransactionService.FindByCustomerId(r.Context(), id)
	if err != nil {
		logger.Error(ctx, logCtx, err)
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerTransactionResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}
