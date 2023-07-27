package app

import (
	"github.com/zanuardi/go-xyz-multifinance/controller"
	"github.com/zanuardi/go-xyz-multifinance/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	customerController controller.CustomerController,
	customerTransactionController controller.CustomerTransactionController,
	customerLimitController controller.CustomerLimitController,
) *httprouter.Router {
	router := httprouter.New()

	// Customer...
	router.GET("/api/customers", customerController.FindAll)
	router.GET("/api/customers/id/:customer_id", customerController.FindById)
	router.POST("/api/customers", customerController.Create)
	router.PUT("/api/customers/id/:customer_id", customerController.UpdateById)
	router.DELETE("/api/customers/id/:customer_id", customerController.DeleteById)

	// Customer Transaction...
	router.GET("/api/customers/transactions/id/:transaction_id", customerTransactionController.FindById)
	router.POST("/api/customers/transactions", customerTransactionController.Create)

	// Customer Limit...
	router.GET("/api/customers/limit", customerLimitController.FindAll)
	router.GET("/api/customers/limit/id/:limit_id", customerLimitController.FindById)
	router.POST("/api/customers/limit", customerLimitController.Create)
	router.PUT("/api/customers/limit/id/:limit_id", customerLimitController.UpdateById)
	router.DELETE("/api/customers/limit/id/:limit_id", customerLimitController.DeleteById)

	router.PanicHandler = exception.ErrorHandler

	return router
}
