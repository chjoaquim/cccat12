package usecase

import (
	"github.com/chjoaquim/ride-service/internal/application/repository"
	"github.com/chjoaquim/ride-service/internal/domain"
)

type CreatePassengerUseCase struct {
	Repository repository.PassengerRepository
}

func NewCreatePassengerUseCase(r repository.PassengerRepository) CreatePassengerUseCase {
	return CreatePassengerUseCase{
		Repository: r,
	}
}

func (s CreatePassengerUseCase) Execute(d *domain.Passenger) (*domain.Passenger, error) {
	result, err := s.Repository.Create(d)
	if err != nil {
		return nil, err
	}
	return result, nil
}
