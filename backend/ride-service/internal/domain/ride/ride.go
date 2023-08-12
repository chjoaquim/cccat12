package ride

import "github.com/chjoaquim/ride-service/internal/domain/ride/segment"

const (
	OvernightFare       = 3.90
	OvernightSundayFare = 5
	SundayFare          = 2.9
	NormalFare          = 2.1
	MinPrice            = 10
)

type Ride struct {
	Segments []segment.Segment `json:"segments"`
}

func (r Ride) Calculate() float64 {
	var price = 0.00

	for _, segment := range r.Segments {
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
	}

	if price < MinPrice {
		price = MinPrice
	}

	return price
}
