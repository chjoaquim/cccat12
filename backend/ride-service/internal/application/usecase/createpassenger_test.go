package usecase

import (
	"errors"
	cpfDomain "github.com/chjoaquim/ride-service/internal/domain/cpf"
	"github.com/chjoaquim/ride-service/internal/domain/passenger"
	"github.com/chjoaquim/ride-service/internal/infra/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGivenAValidPassenger_WhenTryToCreate_ThenReturnPassenger(t *testing.T) {
	cpf, _ := cpfDomain.New("415.765.112-00")
	passenger := &passenger.Passenger{
		ID:        uuid.New().String(),
		Name:      "João",
		Email:     "joão@gmail.com",
		Document:  cpf,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	repository := new(mocks.PassengerRepository)
	repository.On("Create", passenger).Return(passenger, nil)
	uc := NewCreatePassengerUseCase(repository)

	result, err := uc.Execute(passenger)
	assert.Nil(t, err)
	assert.Equal(t, passenger, result)
}

func TestGivenAValidPassenger_WhenTryToCreateWithError_ThenReturnError(t *testing.T) {
	cpf, _ := cpfDomain.New("415.765.112-00")
	passenger := &passenger.Passenger{
		ID:        uuid.New().String(),
		Name:      "João",
		Email:     "joão@gmail.com",
		Document:  cpf,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	repository := new(mocks.PassengerRepository)
	repository.On("Create", passenger).Return(nil, errors.New("test_error"))
	uc := NewCreatePassengerUseCase(repository)

	result, err := uc.Execute(passenger)
	assert.Nil(t, result)
	assert.Equal(t, "test_error", err.Error())
}
