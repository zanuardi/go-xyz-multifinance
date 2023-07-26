package main

import (
	"fmt"
	"net/http"

	"github.com/zanuardi/go-xyz-multifinance/app"
	"github.com/zanuardi/go-xyz-multifinance/controller"
	"github.com/zanuardi/go-xyz-multifinance/helper"
	"github.com/zanuardi/go-xyz-multifinance/logger"
	"github.com/zanuardi/go-xyz-multifinance/middleware"
	"github.com/zanuardi/go-xyz-multifinance/repository"
	"github.com/zanuardi/go-xyz-multifinance/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger.Init("debug")

	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db, validate)
	customerController := controller.NewCustomerController(customerService)

	customerTransactionRepository := repository.NewCustomerRepository()
	customerTransactionService := service.NewCustomerService(customerTransactionRepository, db, validate)
	customerTransactionController := controller.NewCustomerController(customerTransactionService)

	customerInstallmentRepository := repository.NewCustomerRepository()
	customerInstallmentService := service.NewCustomerService(customerInstallmentRepository, db, validate)
	customerInstallmentController := controller.NewCustomerController(customerInstallmentService)

	customerLimitRepository := repository.NewCustomerRepository()
	customerLimitService := service.NewCustomerService(customerLimitRepository, db, validate)
	customerLimitController := controller.NewCustomerController(customerLimitService)

	router := app.NewRouter(
		categoryController,
		customerController,
		customerTransactionController,
		customerInstallmentController,
		customerLimitController,
	)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	fmt.Println("running in", server.Addr)

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
