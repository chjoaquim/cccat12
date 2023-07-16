package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenDateTimeWhenIsOvernightThenReturnTrue(t *testing.T) {
	// Given
	segment := Segment{Date: "2020-01-01T22:00:00Z"}
	// When
	isOvernight := segment.IsOvernight()
	// Then
	assert.True(t, isOvernight)
}

func TestGivenDateTimeWhenIsInvalidThenOvernightReturnFalse(t *testing.T) {
	// Given
	segment := Segment{Date: "invalid"}
	// When
	isOvernight := segment.IsOvernight()
	// Then
	assert.False(t, isOvernight)
}

func TestGivenDateTimeWhenIsNotOvernightThenReturnFalse(t *testing.T) {
	// Given
	segment := Segment{Date: "2020-01-01T10:00:00Z"}
	// When
	isOvernight := segment.IsOvernight()
	// Then
	assert.False(t, isOvernight)
}

func TestGivenDateTimeWhenIsSundayThenReturnTrue(t *testing.T) {
	// Given
	segment := Segment{Date: "2020-01-05T10:00:00Z"}
	// When
	isSunday := segment.IsSunday()
	// Then
	assert.True(t, isSunday)
}

func TestGivenDateTimeWhenIsNotSundayThenReturnFalse(t *testing.T) {
	// Given
	segment := Segment{Date: "2023-07-03T10:00:00Z"}
	// When
	isSunday := segment.IsSunday()
	// Then
	assert.False(t, isSunday)
}

func TestGivenInvalidDateTimeWhenIsNotSundayThenReturnFalse(t *testing.T) {
	// Given
	segment := Segment{Date: "invalid"}
	// When
	isSunday := segment.IsSunday()
	// Then
	assert.False(t, isSunday)
}

func TestGivenDateTimeWhenIsInvalidThenReturnFalse(t *testing.T) {
	// Given
	segment := Segment{Date: "invalid_date"}
	// When
	isValidDate := segment.IsValidDate()
	// Then
	assert.False(t, isValidDate)
}

func TestGivenDateTimeWhenIsValidThenReturnTrue(t *testing.T) {
	// Given
	segment := Segment{Date: "2020-01-01T10:00:00Z"}
	// When
	isValidDate := segment.IsValidDate()
	// Then
	assert.True(t, isValidDate)
}

func TestGivenDistanceWhenValidThenReturnTrue(t *testing.T) {
	// Given
	segment := Segment{Distance: 100}
	// When
	isValidDistance := segment.IsValidDistance()
	// Then
	assert.True(t, isValidDistance)
}

func TestGivenDistanceWhenInValidThenReturnFalse(t *testing.T) {
	// Given
	segment := Segment{Distance: -1}
	// When
	isValidDistance := segment.IsValidDistance()
	// Then
	assert.False(t, isValidDistance)
}

func TestGivenDistanceWhenIsInvalidThenIsValidReturnFalse(t *testing.T) {
	// Given
	segment := Segment{Distance: -1}
	// When
	isValid, err := segment.IsValid()
	// Then
	assert.False(t, isValid)
	assert.Equal(t, "invalid_date", err.Error())
}

func TestGivenDateWhenIsInvalidThenIsValidReturnFalse(t *testing.T) {
	// Given
	segment := Segment{Distance: 10, Date: "invalid"}
	// When
	isValid, err := segment.IsValid()
	// Then
	assert.False(t, isValid)
	assert.Equal(t, "invalid_date", err.Error())
}
