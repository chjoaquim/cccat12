package driver

import (
	"github.com/chjoaquim/ride-service/internal/domain/driver/carplate"
	"github.com/google/uuid"
)

type Driver struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	Document  string            `json:"document"`
	CarPlate  carplate.CarPlate `json:"car_plate"`
	CreatedAt string            `json:"created_at"`
}

func New(name string, email string, document string, carPlate string) (*Driver, error) {
	cp, err := carplate.New(carPlate)
	if err != nil {
		return nil, err
	}
	return &Driver{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Document:  document,
		CarPlate:  cp,
		CreatedAt: "",
	}, nil
}
