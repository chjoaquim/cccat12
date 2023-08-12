package usecase

import (
	"errors"
	"github.com/chjoaquim/ride-service/internal/domain/driver"
	"github.com/chjoaquim/ride-service/internal/infra/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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
	return &driver.Driver{
		ID:        uuid.New().String(),
		Name:      "João",
		Email:     "joão@gmail.com",
		Document:  "123456789",
		CarPlate:  "ABC-1234",
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}
