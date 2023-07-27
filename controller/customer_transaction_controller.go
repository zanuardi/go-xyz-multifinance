package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CustomerTransactionController interface {
	Create(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	FindByCustomerId(w http.ResponseWriter, r *http.Request, param httprouter.Params)
}
