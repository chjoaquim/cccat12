package createpassengers

import (
	"encoding/json"
	"github.com/chjoaquim/ride-service/internal/passengers/domain"
	passengers "github.com/chjoaquim/ride-service/internal/passengers/service"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

type handler struct {
	service passengers.PassengerService
}

func NewCreatePassengersHandler(service passengers.PassengerService) handler {
	return handler{
		service: service,
	}
}

func (h handler) Method() string {
	return http.MethodPost
}

func (h handler) Pattern() string {
	return "/passengers"
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	request := domain.CreatePassenger{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	passenger := request.ToDomain()
	result, err := h.service.Create(&passenger)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}
