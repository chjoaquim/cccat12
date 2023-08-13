package infra

import (
	"fmt"
	cpfDomain "github.com/chjoaquim/ride-service/internal/domain/cpf"
	passengerDomain "github.com/chjoaquim/ride-service/internal/domain/passenger"
	"github.com/chjoaquim/ride-service/pkg/database"
)

type PassengersDB struct {
	db *database.Database
}

func NewPassengersDB(db *database.Database) *PassengersDB {
	return &PassengersDB{
		db: db,
	}
}

func (p *PassengersDB) Get(id string) (*passengerDomain.Passenger, error) {
	passenger := passengerDomain.Passenger{}
	var cpf string
	row := p.db.Connection.QueryRow(`SELECT id, name, document, email, created_at FROM passengers WHERE id=$1`, id)
	err := row.Scan(&passenger.ID, &passenger.Name, &cpf, &passenger.Email, &passenger.CreatedAt)
	if err != nil {
		return nil, err
	}
	c, _ := cpfDomain.New(cpf)
	passenger.Document = c
	return &passenger, nil
}

func (p *PassengersDB) Create(passenger *passengerDomain.Passenger) (*passengerDomain.Passenger, error) {
	stmt, err := p.db.Connection.Prepare(`insert into passengers (id, name, document, email, created_at) values ($1,$2, $3, $4, $5)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		passenger.ID,
		passenger.Name,
		passenger.Document.Value(),
		passenger.Email,
		passenger.CreatedAt,
	)

	if err != nil {
		fmt.Println("Error when trying to create a new passenger: ", err.Error())
		return nil, err
	}

	return passenger, nil
}
