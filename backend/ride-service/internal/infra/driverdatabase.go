package infra

import (
	driverDomain "github.com/chjoaquim/ride-service/internal/domain/driver"
	carPlateDomain "github.com/chjoaquim/ride-service/internal/domain/driver/carplate"
	"github.com/chjoaquim/ride-service/pkg/database"
)

type DriverDB struct {
	db *database.Database
}

func NewDriverDB(db *database.Database) *DriverDB {
	return &DriverDB{
		db: db,
	}
}

func (d *DriverDB) Get(id string) (*driverDomain.Driver, error) {
	driver := driverDomain.Driver{}
	var plate string
	row := d.db.Connection.QueryRow(`SELECT id, name, document, email, car_plate, created_at FROM drivers WHERE id=$1`, id)
	err := row.Scan(&driver.ID, &driver.Name, &driver.Document, &driver.Email, &plate, &driver.CreatedAt)
	if err != nil {
		return nil, err
	}

	carPlate, _ := carPlateDomain.New(plate)
	driver.CarPlate = carPlate
	return &driver, nil
}

func (d *DriverDB) Create(driver *driverDomain.Driver) (*driverDomain.Driver, error) {
	stmt, err := d.db.Connection.Prepare(`insert into drivers (id, name, document, email, car_plate, created_at) values ($1,$2, $3, $4, $5, $6)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		driver.ID,
		driver.Name,
		driver.Document,
		driver.Email,
		driver.CarPlate.Value(),
		driver.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return driver, nil
}
