package repository

import (
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/chjoaquim/ride-service/internal/drivers/domain"
	"github.com/chjoaquim/ride-service/pkg/database"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGivenValidDriver_WhenTryingToInsert_ThenReturnDriver(t *testing.T) {
	db, mock, _ := sqlmock.New()
	database := database.Database{
		Connection: db,
	}
	defer db.Close()
	driverDB := NewDriverDB(&database)
	driver := buildDriver()

	mock.ExpectPrepare("insert into drivers")
	mock.
		ExpectExec("insert into drivers").
		WithArgs(driver.ID, driver.Name, driver.Document, driver.Email, driver.CarPlate, driver.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectBegin()
	mock.ExpectCommit()

	d, err := driverDB.Create(driver)
	assert.Nil(t, err)
	assert.Equal(t, driver.ID, d.ID)
}

func TestGivenValidDriver_WhenTryToPrepareQueryWithError_ThenReturnError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	database := database.Database{
		Connection: db,
	}
	defer db.Close()
	driverDB := NewDriverDB(&database)
	driver := buildDriver()
	mock.ExpectPrepare("insert into drivers").WillReturnError(errors.New("error_to_prepare"))

	p, err := driverDB.Create(driver)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, "error_to_prepare", err.Error())
}

func TestGivenValidDriver_WhenTryToExecQueryWithError_ThenReturnError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	database := database.Database{
		Connection: db,
	}
	defer db.Close()
	driverDB := NewDriverDB(&database)
	driver := buildDriver()
	mock.ExpectPrepare("insert into drivers")
	mock.
		ExpectExec("insert into drivers").
		WithArgs(driver.ID, driver.Name, driver.Document, driver.Email, driver.CarPlate, driver.CreatedAt).
		WillReturnError(errors.New("error_to_exec"))

	p, err := driverDB.Create(driver)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, "error_to_exec", err.Error())
}

func TestGivenValidDriverID_WhenTryGet_ThenReturnOK(t *testing.T) {
	db, mock, _ := sqlmock.New()
	database := database.Database{
		Connection: db,
	}
	defer db.Close()
	driverDB := NewDriverDB(&database)
	driverID := uuid.New().String()
	columns := []string{"id", "name", "document", "email", "car_plate", "created_at"}
	mock.
		ExpectQuery("SELECT id, name, document, email, car_plate, created_at FROM drivers").
		WithArgs(driverID).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString(fmt.Sprintf("%s, Paulo, 13301293, pauloemail.com, AAB-0921, %s", driverID, time.Now().Format(time.RFC3339))))
	mock.ExpectBegin()
	mock.ExpectCommit()

	driver, err := driverDB.Get(driverID)
	assert.Equal(t, driverID, driver.ID)
	assert.Equal(t, "Paulo", driver.Name)
	assert.Equal(t, "13301293", driver.Document)
	assert.Equal(t, "pauloemail.com", driver.Email)
	assert.Equal(t, "AAB-0921", driver.CarPlate)
	assert.Nil(t, err)
}

func TestGivenValidDriverID_WhenGetThrowsAnError_ThenReturnError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	database := database.Database{
		Connection: db,
	}
	defer db.Close()
	driverDB := NewDriverDB(&database)
	driverID := uuid.New().String()
	mock.
		ExpectQuery("SELECT id, name, document, email, car_plate, created_at FROM drivers").
		WithArgs(driverID).
		WillReturnError(errors.New("error_to_get"))
	mock.ExpectBegin()
	mock.ExpectCommit()

	p, err := driverDB.Get(driverID)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, "error_to_get", err.Error())
}

func buildDriver() *domain.Driver {
	return &domain.Driver{
		ID:        uuid.New().String(),
		Name:      "Driver Test",
		CarPlate:  "ABC-1234",
		Document:  "123456789",
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}
