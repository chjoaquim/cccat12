package handlers

import (
	"encoding/json"
	"github.com/chjoaquim/ride-service/internal/drivers/domain"
	"github.com/chjoaquim/ride-service/internal/drivers/services"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

type handler struct {
	service services.DriverService
}

func NewDriverHandler(service services.DriverService) handler {
	return handler{
		service: service,
	}
}

func (h handler) Method() string {
	return http.MethodPost
}

func (h handler) Pattern() string {
	return "/drivers"
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	request := domain.DriverRequest{}
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

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, result)
}
