package usecase

import (
	"errors"
	"github.com/chjoaquim/ride-service/internal/infra/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenValidID_WhenTryToGetDriver_ThenReturnDriver(t *testing.T) {
	driverID := uuid.New().String()
	driver := buildDriver()
	repository := new(mocks.DriverRepository)
	repository.On("Get", driverID).Return(driver, nil)
	getDriver := NewGetDriverUseCase(repository)

	result, err := getDriver.Execute(driverID)
	assert.Nil(t, err)
	assert.Equal(t, driver.Name, result.Name)
}

func TestGivenValidID_WhenTryToGetDriverThrowsAnError_ThenReturnError(t *testing.T) {
	driverID := uuid.New().String()
	repository := new(mocks.DriverRepository)
	repository.On("Get", driverID).Return(nil, errors.New("test_error"))
	getDriver := NewGetDriverUseCase(repository)

	result, err := getDriver.Execute(driverID)
	assert.Nil(t, result)
	assert.Equal(t, "test_error", err.Error())
}
