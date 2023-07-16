package usecase

import "github.com/chjoaquim/ride-service/internal/domain"

type CalculateRide struct {
}

func NewCalculateRide() CalculateRide {
	return CalculateRide{}
}

func (s CalculateRide) Execute(ride domain.Ride) (float64, error) {
	for _, segment := range ride.Segments {
		_, err := segment.IsValid()
		if err != nil {
			return 0.00, err
		}
	}
	return ride.Calculate(), nil
}
