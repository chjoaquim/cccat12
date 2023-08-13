package driver

import (
	cpfDomain "github.com/chjoaquim/ride-service/internal/domain/cpf"
	"github.com/chjoaquim/ride-service/internal/domain/driver/carplate"
	"github.com/google/uuid"
	"time"
)

type Driver struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	Document  cpfDomain.Cpf     `json:"document"`
	CarPlate  carplate.CarPlate `json:"car_plate"`
	CreatedAt string            `json:"created_at"`
}

func New(name string, email string, document string, carPlate string) (*Driver, error) {
	cp, err := carplate.New(carPlate)
	if err != nil {
		return nil, err
	}
	cpf, err := cpfDomain.New(document)
	if err != nil {
		return nil, err
	}

	return &Driver{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Document:  cpf,
		CarPlate:  cp,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}
