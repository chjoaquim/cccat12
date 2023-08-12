package repository

import (
	"github.com/chjoaquim/ride-service/internal/domain/passenger"
)

type PassengerRepository interface {
	Create(passenger *passenger.Passenger) (*passenger.Passenger, error)
	Get(id string) (*passenger.Passenger, error)
}
