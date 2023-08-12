package getdrivers

import (
	"encoding/json"
	"errors"
	"fmt"
	handlermock "github.com/chjoaquim/ride-service/api/mocks"
	"github.com/chjoaquim/ride-service/internal/application/usecase"
	"github.com/chjoaquim/ride-service/internal/domain/driver"
	"github.com/chjoaquim/ride-service/internal/infra/mocks"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGivenValidID_WhenTryTGetDriver_ThenReturnOK(t *testing.T) {

	driver, _ := driver.New("João", "jao@gmail.com", "41565245896", "AAA1234")

	id := driver.ID
	repo := new(mocks.DriverRepository)
	repo.On("Get", id).Return(driver, nil)
	uc := usecase.NewGetDriverUseCase(repo)
	response := sendRequest(id, uc)
	bodyResp := extractBody(response)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "João", bodyResp.Name)
}

func TestGivenInValidRequest_WhenTryToInsertPassenger_ThenReturnBadRequest(t *testing.T) {
	repo := new(mocks.DriverRepository)
	uc := usecase.NewGetDriverUseCase(repo)

	handler := NewGetDriverHandler(uc)
	reader := handlermock.ErrReader(1)
	req, _ := http.NewRequest(handler.Method(), handler.Pattern(), reader)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGivenValidRequest_WhenGetDatabaseError_ThenReturnInternalServerError(t *testing.T) {
	id := "1"
	repo := new(mocks.DriverRepository)
	uc := usecase.NewGetDriverUseCase(repo)
	repo.On("Get", id).Return(nil, errors.New("database error"))

	response := sendRequest(id, uc)
	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func extractBody(response *httptest.ResponseRecorder) *driver.Driver {
	bodyResp := driver.Driver{}
	result, _ := io.ReadAll(response.Body)
	json.Unmarshal(result, &bodyResp)
	return &bodyResp
}

func sendRequest(id string, uc usecase.GetDriverUseCase) *httptest.ResponseRecorder {
	handler := NewGetDriverHandler(uc)
	var request *http.Request
	recorder := httptest.NewRecorder()
	request = httptest.NewRequest(handler.Method(), fmt.Sprintf("/drivers/%s", id), nil)
	mux := chi.NewMux()
	mux.Method(handler.Method(), handler.Pattern(), handler)
	mux.ServeHTTP(recorder, request)
	return recorder
}
