package createpassengers

import (
	"encoding/json"
	"errors"
	"github.com/chjoaquim/ride-service/internal/commons"
	"github.com/chjoaquim/ride-service/internal/passengers/domain"
	handlermock "github.com/chjoaquim/ride-service/internal/passengers/handlers/mocks"
	"github.com/chjoaquim/ride-service/internal/passengers/repository/mocks"
	passengers "github.com/chjoaquim/ride-service/internal/passengers/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGivenValidRequest_WhenTryToInsertPassenger_ThenReturnOK(t *testing.T) {
	request := domain.CreatePassenger{
		Name:     "João",
		Email:    "joao@gmail.com",
		Document: "41565245896",
	}
	passenger := request.ToDomain()
	repo := new(mocks.PassengerRepository)
	repo.On("Create", mock.Anything).Return(&passenger, nil)
	service := passengers.PassengerService{
		Repository: repo,
	}

	response := sendRequest(strings.NewReader(commons.StructToJson(request)), service)
	bodyResp := extractBody(response)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, request.Name, bodyResp.Name)
}

func TestGivenInValidRequest_WhenTryToInsertPassenger_ThenReturnBadRequest(t *testing.T) {
	repo := new(mocks.PassengerRepository)
	service := passengers.PassengerService{
		Repository: repo,
	}

	handler := NewCreatePassengersHandler(service)
	reader := handlermock.ErrReader(1)
	req, _ := http.NewRequest(handler.Method(), handler.Pattern(), reader)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGivenInValidBody_WhenTryToInsertPassenger_ThenReturnBadRequest(t *testing.T) {
	repo := new(mocks.PassengerRepository)
	service := passengers.PassengerService{
		Repository: repo,
	}

	response := sendRequest(strings.NewReader(""), service)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestGivenValidRequest_WhenGetDatabaseError_ThenReturnInternalServerError(t *testing.T) {
	request := domain.CreatePassenger{
		Name:     "João",
		Email:    "joao@gmail.com",
		Document: "41565245896",
	}
	repo := new(mocks.PassengerRepository)
	repo.On("Create", mock.Anything).Return(nil, errors.New("database error"))
	service := passengers.PassengerService{
		Repository: repo,
	}

	response := sendRequest(strings.NewReader(commons.StructToJson(request)), service)
	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func extractBody(response *httptest.ResponseRecorder) *domain.Passenger {
	bodyResp := domain.Passenger{}
	result, _ := io.ReadAll(response.Body)
	json.Unmarshal(result, &bodyResp)
	return &bodyResp
}

func sendRequest(body io.Reader, service passengers.PassengerService) *httptest.ResponseRecorder {
	handler := NewCreatePassengersHandler(service)
	req, _ := http.NewRequest(handler.Method(), handler.Pattern(), body)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}
