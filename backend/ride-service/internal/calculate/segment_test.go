package calculate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenDateTimeWhenIsOvernightThenReturnTrue(t *testing.T) {
	// Given
	segment := Segment{Date: "2020-01-01T22:00:00Z"}
	// When
	isOvernight := segment.isOvernight()
	// Then
	assert.True(t, isOvernight)
}

func TestGivenDateTimeWhenIsInvalidThenOvernightReturnFalse(t *testing.T) {
	// Given
	segment := Segment{Date: "invalid"}
	// When
	isOvernight := segment.isOvernight()
	// Then
	assert.False(t, isOvernight)
}

func TestGivenDateTimeWhenIsNotOvernightThenReturnFalse(t *testing.T) {
	// Given
	segment := Segment{Date: "2020-01-01T10:00:00Z"}
	// When
	isOvernight := segment.isOvernight()
	// Then
	assert.False(t, isOvernight)
}

func TestGivenDateTimeWhenIsSundayThenReturnTrue(t *testing.T) {
	// Given
	segment := Segment{Date: "2020-01-05T10:00:00Z"}
	// When
	isSunday := segment.isSunday()
	// Then
	assert.True(t, isSunday)
}

func TestGivenDateTimeWhenIsNotSundayThenReturnFalse(t *testing.T) {
	// Given
	segment := Segment{Date: "2023-07-03T10:00:00Z"}
	// When
	isSunday := segment.isSunday()
	// Then
	assert.False(t, isSunday)
}

func TestGivenInvalidDateTimeWhenIsNotSundayThenReturnFalse(t *testing.T) {
	// Given
	segment := Segment{Date: "invalid"}
	// When
	isSunday := segment.isSunday()
	// Then
	assert.False(t, isSunday)
}

func TestGivenDateTimeWhenIsInvalidThenReturnFalse(t *testing.T) {
	// Given
	segment := Segment{Date: "invalid_date"}
	// When
	isValidDate := segment.isValidDate()
	// Then
	assert.False(t, isValidDate)
}

func TestGivenDateTimeWhenIsValidThenReturnTrue(t *testing.T) {
	// Given
	segment := Segment{Date: "2020-01-01T10:00:00Z"}
	// When
	isValidDate := segment.isValidDate()
	// Then
	assert.True(t, isValidDate)
}

func TestGivenDistanceWhenValidThenReturnTrue(t *testing.T) {
	// Given
	segment := Segment{Distance: 100}
	// When
	isValidDistance := segment.isValidDistance()
	// Then
	assert.True(t, isValidDistance)
}

func TestGivenDistanceWhenInValidThenReturnFalse(t *testing.T) {
	// Given
	segment := Segment{Distance: -1}
	// When
	isValidDistance := segment.isValidDistance()
	// Then
	assert.False(t, isValidDistance)
}

func TestGivenDistanceWhenIsInvalidThenIsValidReturnFalse(t *testing.T) {
	// Given
	segment := Segment{Distance: -1}
	// When
	isValid, err := segment.isValid()
	// Then
	assert.False(t, isValid)
	assert.Equal(t, "invalid", err.Error())
}

func TestGivenDateWhenIsInvalidThenIsValidReturnFalse(t *testing.T) {
	// Given
	segment := Segment{Distance: 10, Date: "invalid"}
	// When
	isValid, err := segment.isValid()
	// Then
	assert.False(t, isValid)
	assert.Equal(t, "invalid_date", err.Error())
}
