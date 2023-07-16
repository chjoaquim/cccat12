package server

import (
	"github.com/chjoaquim/ride-service/api/createpassengers"
	"github.com/chjoaquim/ride-service/internal/application/repository"
	"github.com/chjoaquim/ride-service/internal/application/usecase"
	"github.com/chjoaquim/ride-service/internal/infra"
	"github.com/chjoaquim/ride-service/pkg/database"
	"go.uber.org/fx"
)

type PassengersHandler struct {
	fx.Out
	Handler HTTPHandler `group:"handlers"`
}

func NewPassengersRepository(db *database.Database) *infra.PassengersDB {
	return infra.NewPassengersDB(db)
}
func NewPassengersUseCase(r repository.PassengerRepository) usecase.CreatePassengerUseCase {
	return usecase.NewCreatePassengerUseCase(r)
}

func NewCreatePassengersHandler(uc usecase.CreatePassengerUseCase) PassengersHandler {
	return PassengersHandler{
		Handler: createpassengers.NewCreatePassengerHandler(uc),
	}
}

var PassengersModule = fx.Provide(
	NewPassengersRepository,
	NewPassengersUseCase,
	NewCreatePassengersHandler,
)
