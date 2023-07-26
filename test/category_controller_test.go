package test

// import (
// 	"context"
// 	"database/sql"
// 	"encoding/json"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"strconv"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/zanuardi/go-xyz-multifinance/app"
// 	"github.com/zanuardi/go-xyz-multifinance/controller"
// 	"github.com/zanuardi/go-xyz-multifinance/helper"
// 	"github.com/zanuardi/go-xyz-multifinance/middleware"
// 	"github.com/zanuardi/go-xyz-multifinance/model/domain"
// 	"github.com/zanuardi/go-xyz-multifinance/repository"
// 	"github.com/zanuardi/go-xyz-multifinance/service"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/stretchr/testify/assert"
// )

// func setupTestDB() *sql.DB {
// 	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/github.com/zanuardi/go-xyz-multifinance_test")
// 	helper.PanicIfError(err)

// 	db.SetMaxIdleConns(5)
// 	db.SetMaxOpenConns(20)
// 	db.SetConnMaxIdleTime(10 * time.Minute)
// 	db.SetConnMaxLifetime(60 * time.Minute)

// 	return db
// }

// func setupRouter(db *sql.DB) http.Handler {
// 	validate := validator.New()
// 	categoryRepository := repository.NewCategoryRepository()
// 	categoryService := service.NewCategoryService(categoryRepository, db, validate)
// 	categoryController := controller.NewCategoryController(categoryService)
// 	router := app.NewRouter(categoryController)

// 	return middleware.NewAuthMiddleware(router)
// }

// func truncateCategory(db *sql.DB) {
// 	db.Exec("TRUNCATE categories")
// }

// func TestCreateCategorySuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)
// 	router := setupRouter(db)

// 	requestBody := strings.NewReader(`{"name":"Makanan"}`)
// 	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-Api-Key", "X-Secret-Key")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)
// }

// func TestCreateCategoryFailed(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)
// 	router := setupRouter(db)

// 	requestBody := strings.NewReader(`{"name":""}`)
// 	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-Api-Key", "X-Secret-Key")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 400, response.StatusCode)

// }

// func TestUpdateCategorySuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	tx, _ := db.Begin()
// 	categoryRepository := repository.NewCategoryRepository()
// 	category, _ := categoryRepository.Create(context.Background(), tx, domain.Category{Name: "Komputer"})
// 	tx.Commit()

// 	router := setupRouter(db)

// 	requestBody := strings.NewReader(`{"name":"Gadget Edit"}`)
// 	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-Api-Key", "X-Secret-Key")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])
// 	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
// 	assert.Equal(t, "Gadget Edit", responseBody["data"].(map[string]interface{})["name"])
// }

// func TestUpdateCategoryFailed(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	tx, _ := db.Begin()
// 	categoryRepository := repository.NewCategoryRepository()
// 	category, _ := categoryRepository.Create(context.Background(), tx, domain.Category{Name: "Komputer"})
// 	tx.Commit()

// 	router := setupRouter(db)

// 	requestBody := strings.NewReader(`{"name":""}`)
// 	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-Api-Key", "X-Secret-Key")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 400, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 400, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "Bad Request.", responseBody["status"])
// }

// func TestGetByIdCategorySuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	tx, _ := db.Begin()
// 	categoryRepository := repository.NewCategoryRepository()
// 	category, _ := categoryRepository.Create(context.Background(), tx, domain.Category{Name: "Komputer"})
// 	tx.Commit()

// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
// 	request.Header.Add("X-Api-Key", "X-Secret-Key")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])
// 	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
// 	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])
// }

// func TestGetByIdCategoryFailed(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)
// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/100", nil)
// 	request.Header.Add("X-Api-Key", "X-Secret-Key")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 404, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 404, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "Data not found.", responseBody["status"])
// }

// func TestDeleteCategorySuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	tx, _ := db.Begin()
// 	categoryRepository := repository.NewCategoryRepository()
// 	category, _ := categoryRepository.Create(context.Background(), tx, domain.Category{Name: "Komputer"})
// 	tx.Commit()

// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-Api-Key", "X-Secret-Key")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])

// }

// func TestDeleteCategoryFailed(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)
// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/555", nil)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-Api-Key", "X-Secret-Key")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 404, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 404, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "Data not found.", responseBody["status"])
// }

// func TestGetAllCategorySuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	tx, _ := db.Begin()
// 	categoryRepository := repository.NewCategoryRepository()
// 	category1, _ := categoryRepository.Create(context.Background(), tx, domain.Category{Name: "Komputer"})
// 	category2, _ := categoryRepository.Create(context.Background(), tx, domain.Category{Name: "Gadget"})

// 	tx.Commit()

// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
// 	request.Header.Add("X-Api-Key", "X-Secret-Key")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])

// 	var categories = responseBody["data"].([]interface{})

// 	categoryResult1 := categories[0].(map[string]interface{})
// 	categoryResult2 := categories[1].(map[string]interface{})

// 	assert.Equal(t, category1.Id, int(categoryResult1["id"].(float64)))
// 	assert.Equal(t, category1.Name, categoryResult1["name"])

// 	assert.Equal(t, category2.Id, int(categoryResult2["id"].(float64)))
// 	assert.Equal(t, category2.Name, categoryResult2["name"])
// }

// func TestUnauthorized(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)
// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/100", nil)
// 	request.Header.Add("X-Api-Key", "Wrong-Secret-Key")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 401, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 401, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "Unauthorized.", responseBody["status"])
// }
