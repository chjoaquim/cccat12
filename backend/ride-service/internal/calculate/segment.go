package calculate

import (
	"errors"
	"time"
)

type Segment struct {
	Distance int64  `json:"distance"`
	Date     string `json:"date"`
}

func (s Segment) isOvernight() bool {
	date, err := time.Parse(time.RFC3339, s.Date)
	if err != nil {
		return false
	}
	return date.Hour() >= 22 || date.Hour() <= 6
}

func (s Segment) isSunday() bool {
	date, err := time.Parse(time.RFC3339, s.Date)
	if err != nil {
		return false
	}
	return date.Weekday() == time.Sunday
}

func (s Segment) isValidDate() bool {
	_, err := time.Parse(time.RFC3339, s.Date)
	return err == nil
}

func (s Segment) isValidDistance() bool {
	return s.Distance > 0
}

func (s Segment) isValid() (bool, error) {
	if !s.isValidDate() {
		return false, errors.New("invalid_date")
	}
	if !s.isValidDistance() {
		return false, errors.New("invalid_distance")
	}
	return true, nil
}
