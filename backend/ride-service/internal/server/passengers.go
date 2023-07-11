package server

import (
	createpassengers "github.com/chjoaquim/ride-service/internal/passengers/handlers"
	"github.com/chjoaquim/ride-service/internal/passengers/repository"
	passengers "github.com/chjoaquim/ride-service/internal/passengers/service"
	"github.com/chjoaquim/ride-service/pkg/database"
	"go.uber.org/fx"
)

type PassengersHandler struct {
	fx.Out
	Handler HTTPHandler `group:"handlers"`
}

func NewPassengersRepository(db *database.Database) repository.PassengerRepository {
	return repository.NewPassengersDb(db)
}
func NewPassengersService(r repository.PassengerRepository) passengers.PassengerService {
	return passengers.NewPassengerService(r)
}

func NewCreatePassengersHandler(passengerService passengers.PassengerService) PassengersHandler {
	return PassengersHandler{
		Handler: createpassengers.NewCreatePassengersHandler(passengerService),
	}
}

var PassengersModule = fx.Provide(
	NewPassengersRepository,
	NewPassengersService,
	NewCreatePassengersHandler,
)
