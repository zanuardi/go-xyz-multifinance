package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CustomerController interface {
	Create(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	UpdateById(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	DeleteById(w http.ResponseWriter, r *http.Request, param httprouter.Params)
}
