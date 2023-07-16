package server

import (
	"github.com/chjoaquim/ride-service/api/createdrivers"
	"github.com/chjoaquim/ride-service/api/getdrivers"
	"github.com/chjoaquim/ride-service/internal/application/repository"
	"github.com/chjoaquim/ride-service/internal/application/usecase"
	"github.com/chjoaquim/ride-service/internal/infra"
	"github.com/chjoaquim/ride-service/pkg/database"
	"go.uber.org/fx"
)

type DriversHandler struct {
	fx.Out
	Handler HTTPHandler `group:"handlers"`
}

func NewDriverRepository(db *database.Database) repository.DriverRepository {
	return infra.NewDriverDB(db)
}
func NewCreateDriversUseCase(r repository.DriverRepository) usecase.CreateDriverUseCase {
	return usecase.NewCreateDriverUseCase(r)
}

func NewGetDriversUseCase(r repository.DriverRepository) usecase.GetDriverUseCase {
	return usecase.NewGetDriverUseCase(r)
}

func NewCreateDriversHandler(uc usecase.CreateDriverUseCase) DriversHandler {
	return DriversHandler{
		Handler: createdrivers.NewCreateDriverHandler(uc),
	}
}

func NewGetDriverHandler(uc usecase.GetDriverUseCase) DriversHandler {
	return DriversHandler{
		Handler: getdrivers.NewGetDriverHandler(uc),
	}
}

var DriversModule = fx.Provide(
	NewDriverRepository,
	NewCreateDriversUseCase,
	NewGetDriversUseCase,
	NewCreateDriversHandler,
	NewGetDriverHandler,
)
