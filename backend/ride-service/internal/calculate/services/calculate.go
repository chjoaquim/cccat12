package services

import "github.com/chjoaquim/ride-service/internal/calculate/domain"

type RideCalculateService struct {
}

func NewRideCalculatorService() RideCalculateService {
	return RideCalculateService{}
}

const (
	OvernightFare       = 3.90
	OvernightSundayFare = 5
	SundayFare          = 2.9
	NormalFare          = 2.1
	MinPrice            = 10
)

func (s RideCalculateService) Calculate(ride domain.Ride) float64 {
	var price = 0.00
	for _, segment := range ride.Segments {
		price += calculateBySegment(segment)
	}

	if price < MinPrice {
		price = MinPrice
	}
	return price
}

func calculateBySegment(segment domain.Segment) float64 {
	var price = 0.00
	if segment.IsOvernight() && !segment.IsSunday() {
		price += float64(segment.Distance) * OvernightFare
	}
	if segment.IsOvernight() && segment.IsSunday() {
		price += float64(segment.Distance) * OvernightSundayFare
	}
	if !segment.IsOvernight() && segment.IsSunday() {
		price += float64(segment.Distance) * SundayFare
	}
	if !segment.IsOvernight() && !segment.IsSunday() {
		price += float64(segment.Distance) * NormalFare
	}
	return price
}