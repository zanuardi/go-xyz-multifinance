package app

import (
	"github.com/zanuardi/go-xyz-multifinance/controller"
	"github.com/zanuardi/go-xyz-multifinance/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:category_id", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:category_id", categoryController.UpdateById)
	router.DELETE("/api/categories/:category_id", categoryController.DeleteById)

	router.PanicHandler = exception.ErrorHandler

	return router
}
