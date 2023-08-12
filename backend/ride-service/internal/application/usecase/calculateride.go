package usecase

import (
	"github.com/chjoaquim/ride-service/internal/domain/ride"
)

type CalculateRide struct {
}

func NewCalculateRide() CalculateRide {
	return CalculateRide{}
}

func (s CalculateRide) Execute(ride ride.Ride) (float64, error) {
	for _, segment := range ride.Segments {
		_, err := segment.IsValid()
		if err != nil {
			return 0.00, err
		}
	}
	return ride.Calculate(), nil
}
