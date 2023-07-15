package domain

type CalculateRequest struct {
	Segments []Segment `json:"segments"`
}

func (cr CalculateRequest) ToDomain() Ride {
	var ride Ride
	for _, segment := range cr.Segments {
		ride.Segments = append(ride.Segments, Segment{
			Distance: segment.Distance,
			Date:     segment.Date,
		})
	}
	return ride
}
