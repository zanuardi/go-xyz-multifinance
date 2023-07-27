package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zanuardi/go-xyz-multifinance/app"
	"github.com/zanuardi/go-xyz-multifinance/controller"
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

	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db, validate)
	customerController := controller.NewCustomerController(customerService)

	customerTransactionRepository := repository.NewCustomerTransactionRepository()
	customerTransactionService := service.NewCustomerTransactionService(customerTransactionRepository, db, validate)
	customerTransactionController := controller.NewCustomerTransactionController(customerTransactionService)

	customerLimitRepository := repository.NewCustomerLimitRepository()
	customerLimitService := service.NewCustomerLimitService(customerLimitRepository, db, validate)
	customerLimitController := controller.NewCustomerLimitController(customerLimitService)

	router := app.NewRouter(
		customerController,
		customerTransactionController,
		customerLimitController,
	)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	fmt.Println("running in", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		logger.Error(context.Background(), "main.server.ListenAndServe", err)
	}
}
