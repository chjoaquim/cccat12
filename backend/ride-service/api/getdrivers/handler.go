package getdrivers

import (
	"github.com/chjoaquim/ride-service/internal/application/usecase"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

type handler struct {
	useCase usecase.GetDriverUseCase
}

func NewGetDriverHandler(uc usecase.GetDriverUseCase) handler {
	return handler{
		useCase: uc,
	}
}

func (h handler) Method() string {
	return http.MethodGet
}

func (h handler) Pattern() string {
	return "/drivers/{id}"
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	result, err := h.useCase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}
