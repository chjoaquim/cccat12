package repository

import (
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/chjoaquim/ride-service/internal/passengers/domain"
	"github.com/chjoaquim/ride-service/pkg/database"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGivenValidPassenger_WhenTryToInsertToDB_ThenReturnOK(t *testing.T) {
	db, mock, _ := sqlmock.New()
	database := database.Database{
		Connection: db,
	}
	defer db.Close()
	pdb := NewPassengersDb(&database)
	passenger := buildPassenger()

	mock.ExpectPrepare("insert into passengers")
	mock.
		ExpectExec("insert into passengers").
		WithArgs(passenger.ID, passenger.Name, passenger.Document, passenger.Email, passenger.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectBegin()
	mock.ExpectCommit()

	p, err := pdb.Create(passenger)
	assert.Nil(t, err)
	assert.NotNil(t, p)
}

func TestGivenValidPassenger_WhenTryToPrepareQueryWithError_ThenReturnError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	database := database.Database{
		Connection: db,
	}
	defer db.Close()
	pdb := NewPassengersDb(&database)
	passenger := buildPassenger()
	mock.ExpectPrepare("insert into passengers").WillReturnError(errors.New("error_to_prepare"))

	p, err := pdb.Create(passenger)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, "error_to_prepare", err.Error())
}

func TestGivenValidPassenger_WhenTryToExecQueryWithError_ThenReturnError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	database := database.Database{
		Connection: db,
	}
	defer db.Close()
	pdb := NewPassengersDb(&database)
	passenger := buildPassenger()
	mock.ExpectPrepare("insert into passengers")
	mock.
		ExpectExec("insert into passengers").
		WithArgs(passenger.ID, passenger.Name, passenger.Document, passenger.Email, passenger.CreatedAt).
		WillReturnError(errors.New("error_to_exec"))

	p, err := pdb.Create(passenger)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, "error_to_exec", err.Error())
}

func TestGivenValidPassengerID_WhenTryGet_ThenReturnOK(t *testing.T) {
	db, mock, _ := sqlmock.New()
	database := database.Database{
		Connection: db,
	}
	defer db.Close()
	pdb := NewPassengersDb(&database)
	passengerID := uuid.New().String()
	columns := []string{"id", "name", "document", "email", "created_at"}
	mock.
		ExpectQuery("SELECT id, name, document, email, created_at FROM passengers").
		WithArgs(passengerID).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString(fmt.Sprintf("%s, João, 13301293, joaoemail.com, %s", passengerID, time.Now().Format(time.RFC3339))))
	mock.ExpectBegin()
	mock.ExpectCommit()

	p, err := pdb.Get(passengerID)
	assert.Equal(t, passengerID, p.ID)
	assert.Equal(t, "João", p.Name)
	assert.Equal(t, "13301293", p.Document)
	assert.Equal(t, "joaoemail.com", p.Email)
	assert.Nil(t, err)
}

func TestGivenValidPassengerID_WhenGetThrowsAnError_ThenReturnError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	database := database.Database{
		Connection: db,
	}
	defer db.Close()
	pdb := NewPassengersDb(&database)
	passengerID := uuid.New().String()
	mock.
		ExpectQuery("SELECT id, name, document, email, created_at FROM passengers").
		WithArgs(passengerID).
		WillReturnError(errors.New("error_to_get"))
	mock.ExpectBegin()
	mock.ExpectCommit()

	p, err := pdb.Get(passengerID)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, "error_to_get", err.Error())
}

func buildPassenger() *domain.Passenger {
	return &domain.Passenger{
		ID:        uuid.New().String(),
		Name:      "João",
		Email:     "joão@gmail.com",
		Document:  "123456789",
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}
