package passenger

import (
	"github.com/chjoaquim/ride-service/internal/domain/cpf"
	"github.com/google/uuid"
	"time"
)

type Passenger struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Document  cpf.Cpf `json:"document"`
	Email     string  `json:"email"`
	CreatedAt string  `json:"created_at"`
}

func New(name, document, email string) (*Passenger, error) {
	cpf, err := cpf.New(document)
	if err != nil {
		return nil, err
	}

	return &Passenger{
		ID:        uuid.New().String(),
		Name:      name,
		Document:  cpf,
		Email:     email,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}
