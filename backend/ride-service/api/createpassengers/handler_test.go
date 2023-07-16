package createpassengers

import (
	"encoding/json"
	"errors"
	handlermock "github.com/chjoaquim/ride-service/api/mocks"
	"github.com/chjoaquim/ride-service/internal/application/usecase"
	"github.com/chjoaquim/ride-service/internal/domain"
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

func TestGivenValidRequest_WhenTryToInsertPassenger_ThenReturnOK(t *testing.T) {
	request := CreatePassenger{
		Name:     "João",
		Email:    "joao@gmail.com",
		Document: "41565245896",
	}
	passenger := request.ToDomain()
	repo := new(mocks.PassengerRepository)
	repo.On("Create", mock.Anything).Return(&passenger, nil)
	uc := usecase.NewCreatePassengerUseCase(repo)
	response := sendRequest(strings.NewReader(commons.StructToJson(request)), uc)
	bodyResp := extractBody(response)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, request.Name, bodyResp.Name)
}

func TestGivenInValidRequest_WhenTryToInsertPassenger_ThenReturnBadRequest(t *testing.T) {
	repo := new(mocks.PassengerRepository)
	uc := usecase.NewCreatePassengerUseCase(repo)

	handler := NewCreatePassengerHandler(uc)
	reader := handlermock.ErrReader(1)
	req, _ := http.NewRequest(handler.Method(), handler.Pattern(), reader)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGivenInValidBody_WhenTryToInsertPassenger_ThenReturnBadRequest(t *testing.T) {
	repo := new(mocks.PassengerRepository)
	uc := usecase.NewCreatePassengerUseCase(repo)

	response := sendRequest(strings.NewReader(""), uc)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestGivenValidRequest_WhenGetDatabaseError_ThenReturnInternalServerError(t *testing.T) {
	request := CreatePassenger{
		Name:     "João",
		Email:    "joao@gmail.com",
		Document: "41565245896",
	}
	repo := new(mocks.PassengerRepository)
	repo.On("Create", mock.Anything).Return(nil, errors.New("database error"))
	uc := usecase.NewCreatePassengerUseCase(repo)

	response := sendRequest(strings.NewReader(commons.StructToJson(request)), uc)
	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func extractBody(response *httptest.ResponseRecorder) *domain.Passenger {
	bodyResp := domain.Passenger{}
	result, _ := io.ReadAll(response.Body)
	json.Unmarshal(result, &bodyResp)
	return &bodyResp
}

func sendRequest(body io.Reader, uc usecase.CreatePassengerUseCase) *httptest.ResponseRecorder {
	handler := NewCreatePassengerHandler(uc)
	req, _ := http.NewRequest(handler.Method(), handler.Pattern(), body)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}
