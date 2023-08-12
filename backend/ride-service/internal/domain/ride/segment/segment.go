package segment

import (
	"errors"
	"time"
)

type Segment struct {
	Distance int64  `json:"distance"`
	Date     string `json:"date"`
}

func (s Segment) IsOvernight() bool {
	date, err := time.Parse(time.RFC3339, s.Date)
	if err != nil {
		return false
	}
	return date.Hour() >= 22 || date.Hour() <= 6
}

func (s Segment) IsSunday() bool {
	date, err := time.Parse(time.RFC3339, s.Date)
	if err != nil {
		return false
	}
	return date.Weekday() == time.Sunday
}

func (s Segment) IsValidDate() bool {
	_, err := time.Parse(time.RFC3339, s.Date)
	return err == nil
}

func (s Segment) IsValidDistance() bool {
	return s.Distance > 0
}

func (s Segment) IsValid() (bool, error) {
	if !s.IsValidDate() {
		return false, errors.New("invalid_date")
	}
	if !s.IsValidDistance() {
		return false, errors.New("invalid_distance")
	}
	return true, nil
}
