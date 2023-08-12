package cpf_test

import (
	"github.com/chjoaquim/ride-service/internal/domain/cpf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenValidCPF_WhenTryToCreate_ThenReturnCPFCreated(t *testing.T) {
	value := "415.765.112-00"
	cpf, err := cpf.New(value)
	assert.Equal(t, value, cpf.Value())
	assert.Nil(t, err)
}

func TestGivenInValidCPF_WhenTryToCreate_ThenReturnError(t *testing.T) {
	value := "135.765.4A8-X0"
	cpf, err := cpf.New(value)
	assert.NotNil(t, err)
	assert.Equal(t, "", cpf.Value())
}

func TestGivenCPFWithLessThanElevenDigits_WhenTryToCreate_ThenReturnError(t *testing.T) {
	value := "123.841.4A8"
	cpf, err := cpf.New(value)
	assert.NotNil(t, err)
	assert.Equal(t, "", cpf.Value())
}

func TestGivenCPFWithMoreThanFourteenDigits_WhenTryToCreate_ThenReturnError(t *testing.T) {
	value := "123.841.4A8.123.123"
	cpf, err := cpf.New(value)
	assert.NotNil(t, err)
	assert.Equal(t, "", cpf.Value())
}

func TestGivenCPFWithOnlyOneDigit_WhenTryToCreate_ThenReturnError(t *testing.T) {
	value := "111.111.111-11"
	cpf, err := cpf.New(value)
	assert.NotNil(t, err)
	assert.Equal(t, "", cpf.Value())
}
