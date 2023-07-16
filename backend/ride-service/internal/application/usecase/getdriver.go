package usecase

import (
	"github.com/chjoaquim/ride-service/internal/application/repository"
	"github.com/chjoaquim/ride-service/internal/domain"
)

type GetDriverUseCase struct {
	Repository repository.DriverRepository
}

func NewGetDriverUseCase(r repository.DriverRepository) GetDriverUseCase {
	return GetDriverUseCase{
		Repository: r,
	}
}

func (s GetDriverUseCase) Execute(id string) (*domain.Driver, error) {
	result, err := s.Repository.Get(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
