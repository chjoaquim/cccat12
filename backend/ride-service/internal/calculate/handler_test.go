package calculate

import (
	"encoding/json"
	"github.com/chjoaquim/ride-service/internal/commons"
	handlermock "github.com/chjoaquim/ride-service/internal/passengers/handlers/mocks"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGivenValidSegment_WhenTryingToCalculate_ThenReturnSuccess(t *testing.T) {
	segment := Segment{
		Distance: 10,
		Date:     "2023-07-11T14:00:00Z",
	}
	service := NewRideCalculatorService()
	response := sendRequest(strings.NewReader(commons.StructToJson(segment)), service)
	result := extractBody(response)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, 21.0, result.Price)
}

func TestGivenInValidSegment_WhenTryingToCalculate_ThenReturnBadRequest(t *testing.T) {
	segment := Segment{
		Distance: -1,
		Date:     "2023-07-11T14:00:00Z",
	}
	service := NewRideCalculatorService()
	response := sendRequest(strings.NewReader(commons.StructToJson(segment)), service)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestGivenInValidSegment_WhenTryToCalculateRide_ThenReturnBadRequest(t *testing.T) {
	service := NewRideCalculatorService()
	response := sendRequest(strings.NewReader(""), service)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestGivenInvalidRequest_WhenTryToCalculateRide_ThenReturnBadRequest(t *testing.T) {
	service := NewRideCalculatorService()
	handler := NewCalculateHandler(service)
	reader := handlermock.ErrReader(1)
	req, _ := http.NewRequest(handler.Method(), handler.Pattern(), reader)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)

}

func sendRequest(body io.Reader, service RideCalculateService) *httptest.ResponseRecorder {
	handler := NewCalculateHandler(service)
	req, _ := http.NewRequest(handler.Method(), handler.Pattern(), body)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func extractBody(response *httptest.ResponseRecorder) *RideCalculateResponse {
	bodyResp := RideCalculateResponse{}
	result, _ := io.ReadAll(response.Body)
	json.Unmarshal(result, &bodyResp)
	return &bodyResp
}