package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/zanuardi/go-xyz-multifinance/app"
	"github.com/zanuardi/go-xyz-multifinance/controller"
	"github.com/zanuardi/go-xyz-multifinance/middleware"
	"github.com/zanuardi/go-xyz-multifinance/model/domain"
	"github.com/zanuardi/go-xyz-multifinance/repository"
	"github.com/zanuardi/go-xyz-multifinance/service"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/xyz_multifinance_test")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
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

	return middleware.NewAuthMiddleware(router)
}

func truncateCustomer(db *sql.DB) {
	db.Exec("TRUNCATE customers;")
}

func TestCreateCustomerSuccess(t *testing.T) {
	db := setupTestDB()
	defer truncateCustomer(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
    		"nik": "2109302139021",
    		"full_name": "Tes Full Name",
			"legal_name": "Tes Legal Name",
			"birth_place": "Tes birth place",
			"birth_date": "1980-10-10T00:00:00Z",
    		"salary": 15000000,
   			"ktp_photo": "url_ktp_photo",
    		"selfie_photo": "url_selfie_photo"
		}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/customers", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "X-Secret-Key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestCreateCustomerFailed(t *testing.T) {
	db := setupTestDB()
	defer truncateCustomer(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/customers", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "X-Secret-Key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 500, response.StatusCode)

}

func TestUpdateCustomerSuccess(t *testing.T) {
	db := setupTestDB()
	defer truncateCustomer(db)

	tx, _ := db.Begin()
	customerRepository := repository.NewCustomerRepository()

	birthDateStr := "1980-10-10T00:00:00Z"
	layout := "2006-01-02T15:04:05Z"
	birthDateTime, _ := time.Parse(layout, birthDateStr)

	customer, _ := customerRepository.Create(context.Background(), tx, domain.Customer{
		Id:          1,
		NIK:         "2109302139021",
		FullName:    "Tes Full Name",
		LegalName:   "Tes Legal Name",
		BirthPlace:  "Tes Birth Place",
		BirthDate:   birthDateTime,
		Salary:      15000000,
		KTPPhoto:    "Tes ktp url",
		SelfiePhoto: "tes selfie photo",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		 "nik": "2109302139021",
         "full_name": "Tes Full Name",
         "legal_name": "Tes Legal Name",
         "birth_place": "Tes Birth Place",
		 "birth_date": "1980-10-10T00:00:00Z",
         "salary": 15000000,
         "ktp_photo": "Tes ktp url",
         "selfie_photo": "tes selfie photo"
		}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/customers/id/"+strconv.Itoa(customer.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "X-Secret-Key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

}

func TestUpdateCustomerFailed(t *testing.T) {
	db := setupTestDB()
	defer truncateCustomer(db)

	tx, _ := db.Begin()
	customerRepository := repository.NewCustomerRepository()

	birthDateStr := "1980-10-10T00:00:00Z"
	layout := "2006-01-02T15:04:05Z"
	birthDateTime, _ := time.Parse(layout, birthDateStr)

	customer, _ := customerRepository.Create(context.Background(), tx, domain.Customer{
		Id:          1,
		NIK:         "2109302139021",
		FullName:    "Tes Full Name",
		LegalName:   "Tes Legal Name",
		BirthPlace:  "Tes Birth Place",
		BirthDate:   birthDateTime,
		Salary:      15000000,
		KTPPhoto:    "Tes ktp url",
		SelfiePhoto: "tes selfie photo"})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"nik":""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/customers/id/"+strconv.Itoa(customer.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "X-Secret-Key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestGetByIdCustomerSuccess(t *testing.T) {
	db := setupTestDB()
	defer truncateCustomer(db)

	tx, _ := db.Begin()
	customerRepository := repository.NewCustomerRepository()

	birthDateStr := "1980-10-10T00:00:00Z"
	layout := "2006-01-02T15:04:05Z"
	birthDateTime, _ := time.Parse(layout, birthDateStr)

	customer, _ := customerRepository.Create(context.Background(), tx, domain.Customer{
		Id:          1,
		NIK:         "2109302139021",
		FullName:    "Tes Full Name",
		LegalName:   "Tes Legal Name",
		BirthPlace:  "Tes Birth Place",
		BirthDate:   birthDateTime,
		Salary:      15000000,
		KTPPhoto:    "Tes ktp url",
		SelfiePhoto: "tes selfie photo"})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/customers/id/"+strconv.Itoa(customer.Id), nil)
	request.Header.Add("X-Api-Key", "X-Secret-Key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

}

func TestGetByIdCustomerFailed(t *testing.T) {
	db := setupTestDB()
	defer truncateCustomer(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/customers/100", nil)
	request.Header.Add("X-Api-Key", "X-Secret-Key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}

func TestDeleteCustomerSuccess(t *testing.T) {
	db := setupTestDB()
	defer truncateCustomer(db)

	tx, _ := db.Begin()
	customerRepository := repository.NewCustomerRepository()

	birthDateStr := "1980-10-10T00:00:00Z"
	layout := "2006-01-02T15:04:05Z"
	birthDateTime, _ := time.Parse(layout, birthDateStr)

	customer, _ := customerRepository.Create(context.Background(), tx, domain.Customer{
		Id:          1,
		NIK:         "2109302139021",
		FullName:    "Tes Full Name",
		LegalName:   "Tes Legal Name",
		BirthPlace:  "Tes Birth Place",
		BirthDate:   birthDateTime,
		Salary:      15000000,
		KTPPhoto:    "Tes ktp url",
		SelfiePhoto: "tes selfie photo",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/customers/id/"+strconv.Itoa(customer.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "X-Secret-Key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

}

func TestDeleteCustomerFailed(t *testing.T) {
	db := setupTestDB()
	defer truncateCustomer(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/customers/id/555", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "X-Secret-Key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}

func TestGetAllCustomerSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)

	tx, _ := db.Begin()
	customerRepository := repository.NewCustomerRepository()

	birthDateStr := "1980-10-10T00:00:00Z"
	layout := "2006-01-02T15:04:05Z"
	birthDateTime, _ := time.Parse(layout, birthDateStr)

	_, _ = customerRepository.Create(context.Background(), tx, domain.Customer{
		Id:          1,
		NIK:         "2109302139021",
		FullName:    "Tes Full Name",
		LegalName:   "Tes Legal Name",
		BirthPlace:  "Tes Birth Place",
		BirthDate:   birthDateTime,
		Salary:      15000000,
		KTPPhoto:    "Tes ktp url",
		SelfiePhoto: "tes selfie photo",
	})
	_, _ = customerRepository.Create(context.Background(), tx, domain.Customer{
		Id:          2,
		NIK:         "2109302139021",
		FullName:    "Tes Full Name",
		LegalName:   "Tes Legal Name",
		BirthPlace:  "Tes Birth Place",
		BirthDate:   birthDateTime,
		Salary:      15000000,
		KTPPhoto:    "Tes ktp url",
		SelfiePhoto: "tes selfie photo",
	})

	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/customers", nil)
	request.Header.Add("X-Api-Key", "X-Secret-Key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/customers/100", nil)
	request.Header.Add("X-Api-Key", "Wrong-Secret-Key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized.", responseBody["status"])
}
