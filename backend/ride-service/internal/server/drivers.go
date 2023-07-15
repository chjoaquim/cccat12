package server

import (
	"github.com/chjoaquim/ride-service/internal/drivers/handlers"
	"github.com/chjoaquim/ride-service/internal/drivers/repository"
	"github.com/chjoaquim/ride-service/internal/drivers/services"
	"github.com/chjoaquim/ride-service/pkg/database"
	"go.uber.org/fx"
)

type DriversHandler struct {
	fx.Out
	Handler HTTPHandler `group:"handlers"`
}

func NewDriverRepository(db *database.Database) repository.DriverRepository {
	return repository.NewDriverDB(db)
}
func NewDriversService(r repository.DriverRepository) services.DriverService {
	return services.NewDriverService(r)
}

func NewCreateDriversHandler(service services.DriverService) DriversHandler {
	return DriversHandler{
		Handler: handlers.NewDriverHandler(service),
	}
}

var DriversModule = fx.Provide(
	NewDriverRepository,
	NewDriversService,
	NewCreateDriversHandler,
)
