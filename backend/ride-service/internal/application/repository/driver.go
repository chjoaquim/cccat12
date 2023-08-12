package repository

import (
	"github.com/chjoaquim/ride-service/internal/domain/driver"
)

type DriverRepository interface {
	Create(driver *driver.Driver) (*driver.Driver, error)
	Get(id string) (*driver.Driver, error)
}
