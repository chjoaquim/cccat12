package usecase

import (
	"github.com/chjoaquim/ride-service/internal/application/repository"
	"github.com/chjoaquim/ride-service/internal/domain/passenger"
)

type CreatePassengerUseCase struct {
	Repository repository.PassengerRepository
}

func NewCreatePassengerUseCase(r repository.PassengerRepository) CreatePassengerUseCase {
	return CreatePassengerUseCase{
		Repository: r,
	}
}

func (s CreatePassengerUseCase) Execute(d *passenger.Passenger) (*passenger.Passenger, error) {
	result, err := s.Repository.Create(d)
	if err != nil {
		return nil, err
	}
	return result, nil
}
