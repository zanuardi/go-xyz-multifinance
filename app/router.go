package app

import (
	"github.com/zanuardi/go-xyz-multifinance/controller"
	"github.com/zanuardi/go-xyz-multifinance/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	categoryController controller.CategoryController,
	customerController controller.CustomerController,
	customerTransactionController controller.CustomerTransactionController,
	customerInstallmentController controller.CustomerInstallmentController,
	customerLimitController controller.CustomerLimitController,
) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:category_id", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:category_id", categoryController.UpdateById)
	router.DELETE("/api/categories/:category_id", categoryController.DeleteById)

	// Customer...
	router.GET("/api/customers", customerController.FindAll)
	router.GET("/api/customers/:customer_id", customerController.FindById)
	router.POST("/api/customers", customerController.Create)
	router.PUT("/api/customers/:customer_id", customerController.UpdateById)
	router.DELETE("/api/customers/:customer_id", customerController.DeleteById)

	// Customer Transaction...
	router.GET("/api/customers/transactions", customerTransactionController.FindAll)
	router.GET("/api/customers/transactions/:transaction_id", customerTransactionController.FindById)
	router.POST("/api/customers/transactions", customerTransactionController.Create)
	router.PUT("/api/customers/transactions/:/transaction_id", customerTransactionController.UpdateById)
	router.DELETE("/api/customers/transactions/:/transaction_id", customerTransactionController.DeleteById)

	// Customer Installment...
	router.GET("/api/customers/installment", customerInstallmentController.FindAll)
	router.GET("/api/customers/installment/:installment_id", customerInstallmentController.FindById)
	router.POST("/api/customers/installment", customerInstallmentController.Create)
	router.PUT("/api/customers/installment/:installment_id", customerInstallmentController.UpdateById)
	router.DELETE("/api/customers/installment/:installent_id", customerInstallmentController.DeleteById)

	// Customer Limit...
	router.GET("/api/customers/limit", customerLimitController.FindAll)
	router.GET("/api/customers/limit/:limit_id", customerLimitController.FindById)
	router.POST("/api/customers/limit", customerLimitController.Create)
	router.PUT("/api/customers/limit/:limit_id", customerLimitController.UpdateById)
	router.DELETE("/api/customers/limit/:limit_id", customerLimitController.DeleteById)

	router.PanicHandler = exception.ErrorHandler

	return router
}
