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

type CustomerTransactionControllerImpl struct {
	customerTransactionService service.CustomerTransactionService
}

func NewCustomerTransactionController(customerTransactionService service.CustomerTransactionService) CustomerTransactionController {
	return &CustomerTransactionControllerImpl{
		customerTransactionService: customerTransactionService,
	}
}

func (customerTransactionController *CustomerTransactionControllerImpl) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	customerTransactionRequest := request.CustomerTransactionRequest{}
	helper.ReadFromRequestBody(r, &customerTransactionRequest)

	customerTransactionResponse, err := customerTransactionController.customerTransactionService.Create(r.Context(), customerTransactionRequest)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerTransactionResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerTransactionController *CustomerTransactionControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	customerTransactionResponses, err := customerTransactionController.customerTransactionService.FindAll(r.Context())
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerTransactionResponses,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerTransactionController *CustomerTransactionControllerImpl) FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	customerTransactionId := param.ByName("transaction_id")
	id, err := strconv.Atoi(customerTransactionId)
	helper.PanicIfError(err)

	customerTransactionResponse, err := customerTransactionController.customerTransactionService.FindById(r.Context(), id)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerTransactionResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerTransactionController *CustomerTransactionControllerImpl) UpdateById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	customerTransactionRequest := request.CustomerTransactionRequest{}
	helper.ReadFromRequestBody(r, &customerTransactionRequest)

	customerTransactionId := param.ByName("transaction_id")
	id, err := strconv.Atoi(customerTransactionId)
	helper.PanicIfError(err)

	customerTransactionRequest.Id = id

	customerTransactionResponse, err := customerTransactionController.customerTransactionService.UpdateById(r.Context(), customerTransactionRequest)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerTransactionResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerTransactionController *CustomerTransactionControllerImpl) DeleteById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	customerTransactionId := param.ByName("transaction_id")
	id, err := strconv.Atoi(customerTransactionId)
	helper.PanicIfError(err)

	err = customerTransactionController.customerTransactionService.DeleteById(r.Context(), id)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)

}
