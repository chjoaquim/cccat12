package usecase

import (
	"github.com/chjoaquim/ride-service/internal/domain/ride"
	"github.com/chjoaquim/ride-service/internal/domain/ride/segment"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenSegmentInOvernight_WhenNotSunday_ThenCalculateWithSuccess(t *testing.T) {
	ride := ride.Ride{
		Segments: []segment.Segment{
			{
				Date:     "2023-07-11T23:00:00Z",
				Distance: 10,
			},
		},
	}

	calculateRide := NewCalculateRide()
	price, err := calculateRide.Execute(ride)
	assert.Equal(t, 39.0, price)
	assert.Nil(t, err)
}

func TestGivenSegmentInOvernight_WhenSunday_ThenCalculateWithSuccess(t *testing.T) {
	ride := ride.Ride{
		Segments: []segment.Segment{
			{
				Date:     "2023-07-16T23:00:00Z",
				Distance: 10,
			},
		},
	}
	calculateRide := NewCalculateRide()
	price, err := calculateRide.Execute(ride)
	assert.Equal(t, 50.0, price)
	assert.Nil(t, err)
}

func TestGivenSegmentNotInOvernight_WhenSunday_ThenCalculateWithSuccess(t *testing.T) {
	ride := ride.Ride{
		Segments: []segment.Segment{
			{
				Date:     "2023-07-16T14:00:00Z",
				Distance: 10,
			},
		},
	}
	calculateRide := NewCalculateRide()
	price, err := calculateRide.Execute(ride)
	assert.Equal(t, 29.0, price)
	assert.Nil(t, err)
}

func TestGivenSegmentNotInOvernight_WhenNotSunday_ThenCalculateWithSuccess(t *testing.T) {
	ride := ride.Ride{
		Segments: []segment.Segment{
			{
				Date:     "2023-07-11T14:00:00Z",
				Distance: 10,
			},
		},
	}
	calculateRide := NewCalculateRide()
	price, err := calculateRide.Execute(ride)
	assert.Equal(t, 21.0, price)
	assert.Nil(t, err)
}

func TestGivenSegment_WhenPriceIsLessThanMinimum_ThenReturnMinimum(t *testing.T) {
	ride := ride.Ride{
		Segments: []segment.Segment{
			{
				Date:     "2023-07-11T14:00:00Z",
				Distance: 1,
			},
		},
	}
	calculateRide := NewCalculateRide()
	price, err := calculateRide.Execute(ride)
	assert.Equal(t, 10.0, price)
	assert.Nil(t, err)
}

func TestGivenInvalidSegment_WhenTryCalculate_ThenReturnError(t *testing.T) {
	ride := ride.Ride{
		Segments: []segment.Segment{
			{
				Date:     "2023-07-11T14:00:00Z",
				Distance: -1,
			},
		},
	}
	calculateRide := NewCalculateRide()
	price, err := calculateRide.Execute(ride)
	assert.Equal(t, 0.00, price)
	assert.Equal(t, "invalid_distance", err.Error())
}
