package carplate_test

import (
	"github.com/chjoaquim/ride-service/internal/domain/driver/carplate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenAValidValueWhenTryToCreateCarPlateThenCreateWithSuccess(t *testing.T) {
	plate := "ABC1234"
	carPlate, err := carplate.New(plate)
	assert.Nil(t, err)
	assert.Equal(t, plate, carPlate.Value())
}

func TestGivenAnInValidValueWhenTryToCreateCarPlateThenReturnAnError(t *testing.T) {
	plate := "AX134"
	carP, err := carplate.New(plate)
	assert.Equal(t, carplate.ErrInvalidCarPlate, err.Error())
	assert.Equal(t, carplate.CarPlate{}, carP)
}

func TestGivenAnInValidRegexWhenTryToCreateCarPlateThenReturnAnError(t *testing.T) {
	plate := "AX134"
	invalidRegex := "xx//\\"
	carplate.RegexPattern = invalidRegex
	carP, err := carplate.New(plate)
	assert.Equal(t, carplate.ErrInvalidCarPlate, err.Error())
	assert.Equal(t, carplate.CarPlate{}, carP)
}
