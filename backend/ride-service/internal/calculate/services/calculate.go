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

func (s RideCalculateService) Calculate(segment domain.Segment) float64 {
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
	if price < MinPrice {
		price = MinPrice
	}
	return price
}
