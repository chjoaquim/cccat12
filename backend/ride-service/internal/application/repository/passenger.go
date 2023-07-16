package repository

import (
	"github.com/chjoaquim/ride-service/internal/domain"
)

type PassengerRepository interface {
	Create(passenger *domain.Passenger) (*domain.Passenger, error)
	Get(id string) (*domain.Passenger, error)
}
