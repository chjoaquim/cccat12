package domain

import (
	"github.com/google/uuid"
	"time"
)

type DriverRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Document string `json:"document"`
	CarPlate string `json:"car_plate"`
}

func (d *DriverRequest) ToDomain() Driver {
	return Driver{
		ID:        uuid.New().String(),
		Name:      d.Name,
		Email:     d.Email,
		Document:  d.Document,
		CarPlate:  d.CarPlate,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}
