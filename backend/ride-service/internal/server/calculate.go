package server

import (
	"github.com/chjoaquim/ride-service/internal/calculate"
	"go.uber.org/fx"
)

type CalculateHandler struct {
	fx.Out
	Handler HTTPHandler `group:"handlers"`
}

func NewCalculateService() calculate.RideCalculateService {
	return calculate.NewRideCalculatorService()
}

func NewCalculateHandler(rideCalculateService calculate.RideCalculateService) CalculateHandler {
	return CalculateHandler{
		Handler: calculate.NewCalculateHandler(rideCalculateService),
	}
}

var CalculateRideModule = fx.Provide(
	NewCalculateService,
	NewCalculateHandler,
)
