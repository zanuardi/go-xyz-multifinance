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

type CustomerLimitControllerImpl struct {
	customerLimitService service.CustomerLimitService
}

func NewCustomerLimitController(customerLimitService service.CustomerLimitService) CustomerLimitController {
	return &CustomerLimitControllerImpl{
		customerLimitService: customerLimitService,
	}
}

func (customerLimitController *CustomerLimitControllerImpl) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	customerLimitRequest := request.CustomerLimitRequest{}
	helper.ReadFromRequestBody(r, &customerLimitRequest)

	customerLimitResponse, err := customerLimitController.customerLimitService.Create(r.Context(), customerLimitRequest)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerLimitResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerLimitController *CustomerLimitControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	customerLimitResponses, err := customerLimitController.customerLimitService.FindAll(r.Context())
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerLimitResponses,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerLimitController *CustomerLimitControllerImpl) FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	customerLimitId := param.ByName("limit_id")
	id, err := strconv.Atoi(customerLimitId)
	helper.PanicIfError(err)

	customerLimitResponse, err := customerLimitController.customerLimitService.FindById(r.Context(), id)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerLimitResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerLimitController *CustomerLimitControllerImpl) UpdateById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	customerLimitRequest := request.CustomerLimitRequest{}
	helper.ReadFromRequestBody(r, &customerLimitRequest)

	customerLimitId := param.ByName("limit_id")
	id, err := strconv.Atoi(customerLimitId)
	helper.PanicIfError(err)

	customerLimitRequest.Id = id

	customerLimitResponse, err := customerLimitController.customerLimitService.UpdateById(r.Context(), customerLimitRequest)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerLimitResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerLimitController *CustomerLimitControllerImpl) DeleteById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	customerLimitId := param.ByName("limit_id")
	id, err := strconv.Atoi(customerLimitId)
	helper.PanicIfError(err)

	err = customerLimitController.customerLimitService.DeleteById(r.Context(), id)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)

}
