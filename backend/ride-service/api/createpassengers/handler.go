package createpassengers

import (
	"encoding/json"
	"github.com/chjoaquim/ride-service/internal/application/usecase"
	"github.com/chjoaquim/ride-service/internal/domain"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"io"
	"net/http"
	"time"
)

type CreatePassenger struct {
	Name     string `json:"name"`
	Document string `json:"document"`
	Email    string `json:"email"`
}

func (c *CreatePassenger) ToDomain() domain.Passenger {
	return domain.Passenger{
		ID:        uuid.New().String(),
		Name:      c.Name,
		Email:     c.Email,
		Document:  c.Document,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
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

	passenger := request.ToDomain()
	result, err := h.CreatePassengerUseCase.Execute(&passenger)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, result)
}
