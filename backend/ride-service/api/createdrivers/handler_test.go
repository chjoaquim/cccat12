package createdrivers

import (
	"encoding/json"
	"errors"
	handlermock "github.com/chjoaquim/ride-service/api/mocks"
	"github.com/chjoaquim/ride-service/internal/application/usecase"
	"github.com/chjoaquim/ride-service/internal/domain/driver"
	"github.com/chjoaquim/ride-service/internal/infra/mocks"
	"github.com/chjoaquim/ride-service/pkg/commons"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGivenValidRequest_WhenTryToInsertDriver_ThenReturnOK(t *testing.T) {
	request := DriverRequest{
		Name:     "João",
		Email:    "joao@gmail.com",
		CarPlate: "AAA1234",
		Document: "41565245896",
	}
	driver, _ := request.ToDomain()
	repo := new(mocks.DriverRepository)
	repo.On("Create", mock.Anything).Return(driver, nil)
	uc := usecase.NewCreateDriverUseCase(repo)

	response := sendRequest(strings.NewReader(commons.StructToJson(request)), uc)
	bodyResp := extractBody(response)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, request.Name, bodyResp.Name)
}

func TestGivenInValidRequest_WhenTryToInsertDriver_ThenReturnBadRequest(t *testing.T) {
	repo := new(mocks.DriverRepository)
	uc := usecase.NewCreateDriverUseCase(repo)
	handler := NewCreateDriverHandler(uc)
	reader := handlermock.ErrReader(1)
	req, _ := http.NewRequest(handler.Method(), handler.Pattern(), reader)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGivenInValidCarPlate_WhenTryToInsertDriver_ThenReturnBadRequest(t *testing.T) {
	request := DriverRequest{
		Name:     "João",
		Email:    "joao@gmail.com",
		CarPlate: "AX04",
		Document: "41565245896",
	}
	driver, _ := request.ToDomain()
	repo := new(mocks.DriverRepository)
	repo.On("Create", mock.Anything).Return(driver, nil)
	uc := usecase.NewCreateDriverUseCase(repo)

	response := sendRequest(strings.NewReader(commons.StructToJson(request)), uc)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestGivenInValidBody_WhenTryToInsertDriver_ThenReturnBadRequest(t *testing.T) {
	repo := new(mocks.DriverRepository)
	uc := usecase.NewCreateDriverUseCase(repo)
	response := sendRequest(strings.NewReader(""), uc)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestGivenValidRequest_WhenGetDatabaseError_ThenReturnInternalServerError(t *testing.T) {
	request := DriverRequest{
		Name:     "João",
		Email:    "joao@gmail.com",
		CarPlate: "AAA1234",
		Document: "41565245896",
	}
	repo := new(mocks.DriverRepository)
	repo.On("Create", mock.Anything).Return(nil, errors.New("database error"))
	uc := usecase.NewCreateDriverUseCase(repo)
	response := sendRequest(strings.NewReader(commons.StructToJson(request)), uc)
	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func sendRequest(body io.Reader, uc usecase.CreateDriverUseCase) *httptest.ResponseRecorder {
	handler := NewCreateDriverHandler(uc)
	req, _ := http.NewRequest(handler.Method(), handler.Pattern(), body)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func extractBody(response *httptest.ResponseRecorder) *driver.Driver {
	bodyResp := driver.Driver{}
	result, _ := io.ReadAll(response.Body)
	json.Unmarshal(result, &bodyResp)
	return &bodyResp
}
