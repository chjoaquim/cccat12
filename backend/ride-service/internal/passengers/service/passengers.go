package passengers

import (
	"fmt"
	"github.com/chjoaquim/ride-service/internal/passengers/domain"
	"github.com/chjoaquim/ride-service/internal/passengers/repository"
)

type PassengerService struct {
	Repository repository.PassengerRepository
}

func NewPassengerService(repository repository.PassengerRepository) PassengerService {
	return PassengerService{
		repository,
	}
}

func (s PassengerService) Create(p *domain.Passenger) (*domain.Passenger, error) {
	result, err := s.Repository.Create(p)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return result, nil
}
