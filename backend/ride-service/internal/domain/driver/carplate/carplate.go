package carplate

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidCarPlate = "invalid_car_plate"
	RegexPattern       = "[A-Z]{3}[0-9]{4}"
)

type CarPlate struct {
	value string
}

func New(value string) (CarPlate, error) {
	value = strings.ToUpper(value)
	if !isValid(value) {
		return CarPlate{}, errors.New(ErrInvalidCarPlate)
	}
	return CarPlate{value: value}, nil
}

func (c CarPlate) Value() string {
	return c.value
}

func isValid(carPlate string) bool {
	isMatched, err := regexp.MatchString(RegexPattern, carPlate)
	if err != nil {
		return false
	}
	return isMatched
}
