package createpassengers

import (
	"encoding/json"
	"github.com/chjoaquim/ride-service/internal/application/usecase"
	"github.com/chjoaquim/ride-service/internal/domain/passenger"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

type CreatePassenger struct {
	Name     string `json:"name"`
	Document string `json:"document"`
	Email    string `json:"email"`
}

func (c *CreatePassenger) ToDomain() (*passenger.Passenger, error) {
	return passenger.New(c.Name, c.Document, c.Email)
}

type handler struct {
	usecase.CreatePassengerUseCase
}

func NewCreatePassengerHandler(uc usecase.CreatePassengerUseCase) handler {
	return handler{
		CreatePassengerUseCase: uc,
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

	request := CreatePassenger{}
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
	result, err := h.CreatePassengerUseCase.Execute(passenger)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, result)
}
