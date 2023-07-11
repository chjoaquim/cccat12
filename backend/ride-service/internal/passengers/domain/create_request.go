package domain

import (
	"github.com/google/uuid"
	"time"
)

type CreatePassenger struct {
	Name     string `json:"name"`
	Document string `json:"document"`
	Email    string `json:"email"`
}

func (c *CreatePassenger) ToDomain() Passenger {
	return Passenger{
		ID:        uuid.New().String(),
		Name:      c.Name,
		Email:     c.Email,
		Document:  c.Document,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}
