package repository

import (
	"fmt"
	"github.com/chjoaquim/ride-service/internal/passengers/domain"
	"github.com/chjoaquim/ride-service/pkg/database"
)

type PassengersDb struct {
	db *database.Database
}

type PassengerRepository interface {
	Create(passenger *domain.Passenger) (*domain.Passenger, error)
	Get(id string) (*domain.Passenger, error)
}

func NewPassengersDb(db *database.Database) *PassengersDb {
	return &PassengersDb{
		db: db,
	}
}

func (p *PassengersDb) Get(id string) (*domain.Passenger, error) {
	passenger := domain.Passenger{}
	row := p.db.Connection.QueryRow(`SELECT id, name, document, email, created_at FROM passengers WHERE id=$1`, id)
	err := row.Scan(&passenger.ID, &passenger.Name, &passenger.Document, &passenger.Email, &passenger.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &passenger, nil
}

func (p *PassengersDb) Create(passenger *domain.Passenger) (*domain.Passenger, error) {
	stmt, err := p.db.Connection.Prepare(`insert into passengers (id, name, document, email, created_at) values ($1,$2, $3, $4, $5)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		passenger.ID,
		passenger.Name,
		passenger.Document,
		passenger.Email,
		passenger.CreatedAt,
	)

	if err != nil {
		fmt.Println("Error when trying to create a new passenger: ", err.Error())
		return nil, err
	}

	return passenger, nil
}
