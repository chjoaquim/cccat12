package repository

import "github.com/chjoaquim/ride-service/internal/domain"

type DriverRepository interface {
	Create(driver *domain.Driver) (*domain.Driver, error)
	Get(id string) (*domain.Driver, error)
}
