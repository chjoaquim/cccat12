package usecase

import (
	"errors"
	"github.com/chjoaquim/ride-service/internal/domain/driver"
	"github.com/chjoaquim/ride-service/internal/infra/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenAValidDriver_WhenTryToCreate_ThenReturnDriver(t *testing.T) {
	driver := buildDriver()
	repository := new(mocks.DriverRepository)
	repository.On("Create", driver).Return(driver, nil)
	createDriver := NewCreateDriverUseCase(repository)

	result, err := createDriver.Execute(driver)
	assert.Nil(t, err)
	assert.Equal(t, driver, result)
}

func TestGivenAValidDriver_WhenTryToCreateWithError_ThenReturnError(t *testing.T) {
	driver := buildDriver()
	repository := new(mocks.DriverRepository)
	repository.On("Create", driver).Return(nil, errors.New("test_error"))
	createDriver := NewCreateDriverUseCase(repository)

	result, err := createDriver.Execute(driver)
	assert.Nil(t, result)
	assert.Equal(t, "test_error", err.Error())
}

func buildDriver() *driver.Driver {
	driver, _ := driver.New("João", "joão@gmail.com", "415.765.112-00", "ABC1234")
	return driver
}
