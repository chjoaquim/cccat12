package services

import (
	"github.com/chjoaquim/ride-service/internal/drivers/domain"
	"github.com/chjoaquim/ride-service/internal/drivers/repository"
)

type DriverService struct {
	Repository repository.DriverRepository
}

func NewDriverService(r repository.DriverRepository) DriverService {
	return DriverService{
		Repository: r,
	}
}

func (s DriverService) Create(d *domain.Driver) (*domain.Driver, error) {
	result, err := s.Repository.Create(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s DriverService) Get(id string) (*domain.Driver, error) {
	result, err := s.Repository.Get(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
