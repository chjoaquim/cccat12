package calculate

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

func (s RideCalculateService) Calculate(segment Segment) float64 {
	var price = 0.00
	if segment.isOvernight() && !segment.isSunday() {
		price += float64(segment.Distance) * OvernightFare
	}
	if segment.isOvernight() && segment.isSunday() {
		price += float64(segment.Distance) * OvernightSundayFare
	}
	if !segment.isOvernight() && segment.isSunday() {
		price += float64(segment.Distance) * SundayFare
	}
	if !segment.isOvernight() && !segment.isSunday() {
		price += float64(segment.Distance) * NormalFare
	}
	if price < MinPrice {
		price = MinPrice
	}
	return price
}
