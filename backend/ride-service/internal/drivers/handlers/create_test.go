package handlers

import (
	"encoding/json"
	"errors"
	"github.com/chjoaquim/ride-service/internal/commons"
	"github.com/chjoaquim/ride-service/internal/drivers/domain"
	"github.com/chjoaquim/ride-service/internal/drivers/repository/mocks"
	"github.com/chjoaquim/ride-service/internal/drivers/services"
	handlermock "github.com/chjoaquim/ride-service/internal/passengers/handlers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGivenValidRequest_WhenTryToInsertDriver_ThenReturnOK(t *testing.T) {
	request := domain.DriverRequest{
		Name:     "João",
		Email:    "joao@gmail.com",
		CarPlate: "AAA-1234",
		Document: "41565245896",
	}
	driver := request.ToDomain()
	repo := new(mocks.DriverRepository)
	repo.On("Create", mock.Anything).Return(&driver, nil)
	service := services.NewDriverService(repo)

	response := sendRequest(strings.NewReader(commons.StructToJson(request)), service)
	bodyResp := extractBody(response)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, request.Name, bodyResp.Name)
}

func TestGivenInValidRequest_WhenTryToInsertDriver_ThenReturnBadRequest(t *testing.T) {
	repo := new(mocks.DriverRepository)
	service := services.NewDriverService(repo)
	handler := NewDriverHandler(service)
	reader := handlermock.ErrReader(1)
	req, _ := http.NewRequest(handler.Method(), handler.Pattern(), reader)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGivenInValidBody_WhenTryToInsertDriver_ThenReturnBadRequest(t *testing.T) {
	repo := new(mocks.DriverRepository)
	service := services.NewDriverService(repo)
	response := sendRequest(strings.NewReader(""), service)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestGivenValidRequest_WhenGetDatabaseError_ThenReturnInternalServerError(t *testing.T) {
	request := domain.DriverRequest{
		Name:     "João",
		Email:    "joao@gmail.com",
		CarPlate: "AAA-1234",
		Document: "41565245896",
	}
	repo := new(mocks.DriverRepository)
	repo.On("Create", mock.Anything).Return(nil, errors.New("database error"))
	service := services.NewDriverService(repo)
	response := sendRequest(strings.NewReader(commons.StructToJson(request)), service)
	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func sendRequest(body io.Reader, service services.DriverService) *httptest.ResponseRecorder {
	handler := NewDriverHandler(service)
	req, _ := http.NewRequest(handler.Method(), handler.Pattern(), body)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func extractBody(response *httptest.ResponseRecorder) *domain.Driver {
	bodyResp := domain.Driver{}
	result, _ := io.ReadAll(response.Body)
	json.Unmarshal(result, &bodyResp)
	return &bodyResp
}
