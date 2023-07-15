package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	Connection *sql.DB
}

func NewDatabase(c *Config) *Database {
	db := Database{}
	db.Connect(c)
	return &db
}

func (d *Database) Connect(c *Config) {
	db, err := sql.Open("postgres", dsn(c))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	d.Connection = db
}

func dsn(c *Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable search_path=%s",
		c.Credential.Host, c.Credential.Port, c.Credential.Username, c.Credential.Password, c.Credential.Name, c.Credential.Schema)
}
