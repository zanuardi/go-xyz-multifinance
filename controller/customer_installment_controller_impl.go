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

type CustomerInstallmentControllerImpl struct {
	customerInstallmentService service.CustomerInstallmentService
}

func NewCustomerInstallmentController(customerInstallmentService service.CustomerInstallmentService) CustomerInstallmentController {
	return &CustomerInstallmentControllerImpl{
		customerInstallmentService: customerInstallmentService,
	}
}

func (customerInstallmentController *CustomerInstallmentControllerImpl) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	customerInstallmentRequest := request.CustomerInstallmentRequest{}
	helper.ReadFromRequestBody(r, &customerInstallmentRequest)

	customerInstallmentResponse, err := customerInstallmentController.customerInstallmentService.Create(r.Context(), customerInstallmentRequest)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerInstallmentResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerInstallmentController *CustomerInstallmentControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	customerInstallmentResponses, err := customerInstallmentController.customerInstallmentService.FindAll(r.Context())
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerInstallmentResponses,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerInstallmentController *CustomerInstallmentControllerImpl) FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	customerInstallmentId := param.ByName("installment_id")
	id, err := strconv.Atoi(customerInstallmentId)
	helper.PanicIfError(err)

	customerInstallmentResponse, err := customerInstallmentController.customerInstallmentService.FindById(r.Context(), id)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerInstallmentResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerInstallmentController *CustomerInstallmentControllerImpl) UpdateById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	customerInstallmentRequest := request.CustomerInstallmentRequest{}
	helper.ReadFromRequestBody(r, &customerInstallmentRequest)

	customerInstallmentId := param.ByName("installment_id")
	id, err := strconv.Atoi(customerInstallmentId)
	helper.PanicIfError(err)

	customerInstallmentRequest.Id = id

	customerInstallmentResponse, err := customerInstallmentController.customerInstallmentService.UpdateById(r.Context(), customerInstallmentRequest)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerInstallmentResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (customerInstallmentController *CustomerInstallmentControllerImpl) DeleteById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	customerInstallmentId := param.ByName("installment_id")
	id, err := strconv.Atoi(customerInstallmentId)
	helper.PanicIfError(err)

	err = customerInstallmentController.customerInstallmentService.DeleteById(r.Context(), id)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)

}
