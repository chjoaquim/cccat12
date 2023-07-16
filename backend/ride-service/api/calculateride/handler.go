package calculateride

import (
	"encoding/json"
	"github.com/chjoaquim/ride-service/internal/application/usecase"
	"github.com/chjoaquim/ride-service/internal/domain"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

type handler struct {
	calculateRide usecase.CalculateRide
}

type RideCalculateResponse struct {
	Price float64 `json:"price"`
}

type CalculateRequest struct {
	Segments []domain.Segment `json:"segments"`
}

func (cr CalculateRequest) ToDomain() domain.Ride {
	var ride domain.Ride
	for _, segment := range cr.Segments {
		ride.Segments = append(ride.Segments, domain.Segment{
			Distance: segment.Distance,
			Date:     segment.Date,
		})
	}
	return ride
}

func NewCalculateHandler(calculateRide usecase.CalculateRide) handler {
	return handler{
		calculateRide: calculateRide,
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

	var request CalculateRequest
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

	result := h.calculateRide.Execute(request.ToDomain())
	response := RideCalculateResponse{
		Price: result,
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}
