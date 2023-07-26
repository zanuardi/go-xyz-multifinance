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

type CustomerControllerImpl struct {
	customerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		customerService: customerService,
	}
}

func (customerController *CustomerControllerImpl) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	customerRequest := request.CustomerRequest{}
	helper.ReadFromRequestBody(r, &customerRequest)

	customerResponse, err := customerController.customerService.Create(r.Context(), customerRequest)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerController *CustomerControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	customerResponses, err := customerController.customerService.FindAll(r.Context())
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponses,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerController *CustomerControllerImpl) FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	customerId := param.ByName("customer_id")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	customerResponse, err := customerController.customerService.FindById(r.Context(), id)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerController *CustomerControllerImpl) UpdateById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	customerRequest := request.CustomerRequest{}
	helper.ReadFromRequestBody(r, &customerRequest)

	customerId := param.ByName("customer_id")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	customerRequest.Id = id

	customerResponse, err := customerController.customerService.UpdateById(r.Context(), customerRequest)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerController *CustomerControllerImpl) DeleteById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	customerId := param.ByName("customer_id")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	err = customerController.customerService.DeleteById(r.Context(), id)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)

}
