package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenSegmentInOvernight_WhenNotSunday_ThenCalculateWithSuccess(t *testing.T) {
	ride := Ride{
		Segments: []Segment{
			{
				Date:     "2023-07-11T23:00:00Z",
				Distance: 10,
			},
		},
	}

	price := ride.Calculate()
	assert.Equal(t, 39.0, price)
}

func TestGivenSegmentInOvernight_WhenSunday_ThenCalculateWithSuccess(t *testing.T) {
	ride := Ride{
		Segments: []Segment{
			{
				Date:     "2023-07-16T23:00:00Z",
				Distance: 10,
			},
		},
	}

	price := ride.Calculate()
	assert.Equal(t, 50.0, price)
}

func TestGivenSegmentNotInOvernight_WhenSunday_ThenCalculateWithSuccess(t *testing.T) {
	ride := Ride{
		Segments: []Segment{
			{
				Date:     "2023-07-16T14:00:00Z",
				Distance: 10,
			},
		},
	}
	price := ride.Calculate()
	assert.Equal(t, 29.0, price)
}

func TestGivenSegmentNotInOvernight_WhenNotSunday_ThenCalculateWithSuccess(t *testing.T) {
	ride := Ride{
		Segments: []Segment{
			{
				Date:     "2023-07-11T14:00:00Z",
				Distance: 10,
			},
		},
	}
	price := ride.Calculate()
	assert.Equal(t, 21.0, price)
}

func TestGivenSegment_WhenPriceIsLessThanMinimum_ThenReturnMinimum(t *testing.T) {
	ride := Ride{
		Segments: []Segment{
			{
				Date:     "2023-07-11T14:00:00Z",
				Distance: 1,
			},
		},
	}
	price := ride.Calculate()
	assert.Equal(t, 10.0, price)
}
