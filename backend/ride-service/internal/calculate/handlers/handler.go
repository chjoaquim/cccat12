package handlers

import (
	"encoding/json"
	"github.com/chjoaquim/ride-service/internal/calculate/domain"
	"github.com/chjoaquim/ride-service/internal/calculate/services"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

type handler struct {
	calculateService services.RideCalculateService
}

type RideCalculateResponse struct {
	Price float64 `json:"price"`
}

func NewCalculateHandler(calculateService services.RideCalculateService) handler {
	return handler{
		calculateService: calculateService,
	}
}

func (h handler) Method() string {
	return http.MethodPost
}

func (h handler) Pattern() string {
	return "/calculate"
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request domain.CalculateRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, segment := range request.Segments {
		_, err = segment.IsValid()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	result := h.calculateService.Calculate(request.ToDomain())
	response := RideCalculateResponse{
		Price: result,
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}