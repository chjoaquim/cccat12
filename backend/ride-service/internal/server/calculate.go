package server

import (
	"github.com/chjoaquim/ride-service/internal/calculate/handlers"
	service "github.com/chjoaquim/ride-service/internal/calculate/services"
	"go.uber.org/fx"
)

type CalculateHandler struct {
	fx.Out
	Handler HTTPHandler `group:"handlers"`
}

func NewCalculateService() service.RideCalculateService {
	return service.NewRideCalculatorService()
}

func NewCalculateHandler(rideCalculateService service.RideCalculateService) CalculateHandler {
	return CalculateHandler{
		Handler: handlers.NewCalculateHandler(rideCalculateService),
	}
}

var CalculateRideModule = fx.Provide(
	NewCalculateService,
	NewCalculateHandler,
)
