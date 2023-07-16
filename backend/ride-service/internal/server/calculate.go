package server

import (
	"github.com/chjoaquim/ride-service/api/calculateride"
	"github.com/chjoaquim/ride-service/internal/application/usecase"
	"go.uber.org/fx"
)

type CalculateHandler struct {
	fx.Out
	Handler HTTPHandler `group:"handlers"`
}

func NewCalculateUseCase() usecase.CalculateRide {
	return usecase.NewCalculateRide()
}

func NewCalculateHandler(uc usecase.CalculateRide) CalculateHandler {
	return CalculateHandler{
		Handler: calculateride.NewCalculateHandler(uc),
	}
}

var CalculateRideModule = fx.Provide(
	NewCalculateUseCase,
	NewCalculateHandler,
)
