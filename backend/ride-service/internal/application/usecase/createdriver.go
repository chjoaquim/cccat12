package usecase

import (
	"github.com/chjoaquim/ride-service/internal/application/repository"
	"github.com/chjoaquim/ride-service/internal/domain/driver"
)

type CreateDriverUseCase struct {
	Repository repository.DriverRepository
}

func NewCreateDriverUseCase(r repository.DriverRepository) CreateDriverUseCase {
	return CreateDriverUseCase{
		Repository: r,
	}
}

func (s CreateDriverUseCase) Execute(d *driver.Driver) (*driver.Driver, error) {
	result, err := s.Repository.Create(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}
