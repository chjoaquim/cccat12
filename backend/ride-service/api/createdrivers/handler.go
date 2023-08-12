package createdrivers

import (
	"encoding/json"
	"github.com/chjoaquim/ride-service/internal/application/usecase"
	"github.com/chjoaquim/ride-service/internal/domain/driver"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

type handler struct {
	useCase usecase.CreateDriverUseCase
}

type DriverRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Document string `json:"document"`
	CarPlate string `json:"car_plate"`
}

func (d *DriverRequest) ToDomain() (*driver.Driver, error) {
	return driver.New(d.Name, d.Email, d.Document, d.CarPlate)
}

func NewCreateDriverHandler(uc usecase.CreateDriverUseCase) handler {
	return handler{
		useCase: uc,
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

	request := DriverRequest{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	passenger, err := request.ToDomain()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.useCase.Execute(passenger)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, result)
}
