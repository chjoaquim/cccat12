package services

import (
	"errors"
	"github.com/chjoaquim/ride-service/internal/drivers/domain"
	"github.com/chjoaquim/ride-service/internal/drivers/repository/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGivenAValidDriver_WhenTryToCreate_ThenReturnDriver(t *testing.T) {
	driver := buildDriver()
	repository := new(mocks.DriverRepository)
	repository.On("Create", driver).Return(driver, nil)
	service := NewDriverService(repository)

	result, err := service.Create(driver)
	assert.Nil(t, err)
	assert.Equal(t, driver, result)
}

func TestGivenAValidDriver_WhenTryToCreateWithError_ThenReturnError(t *testing.T) {
	driver := buildDriver()
	repository := new(mocks.DriverRepository)
	repository.On("Create", driver).Return(nil, errors.New("test_error"))
	service := NewDriverService(repository)

	result, err := service.Create(driver)
	assert.Nil(t, result)
	assert.Equal(t, "test_error", err.Error())
}

func TestGivenValidID_WhenTryToGetDriver_ThenReturnDriver(t *testing.T) {
	driverID := uuid.New().String()
	driver := buildDriver()
	repository := new(mocks.DriverRepository)
	repository.On("Get", driverID).Return(driver, nil)
	service := NewDriverService(repository)

	result, err := service.Get(driverID)
	assert.Nil(t, err)
	assert.Equal(t, driver.Name, result.Name)
}

func TestGivenValidID_WhenTryToGetDriverThrowsAnError_ThenReturnError(t *testing.T) {
	driverID := uuid.New().String()
	repository := new(mocks.DriverRepository)
	repository.On("Get", driverID).Return(nil, errors.New("test_error"))
	service := NewDriverService(repository)

	result, err := service.Get(driverID)
	assert.Nil(t, result)
	assert.Equal(t, "test_error", err.Error())
}

func buildDriver() *domain.Driver {
	return &domain.Driver{
		ID:        uuid.New().String(),
		Name:      "João",
		Email:     "joão@gmail.com",
		Document:  "123456789",
		CarPlate:  "ABC-1234",
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}
