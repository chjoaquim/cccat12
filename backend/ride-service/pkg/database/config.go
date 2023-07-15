package database

type Config struct {
	Credential Credential
}

type Credential struct {
	Username string
	Password string
	Host     string
	Name     string
	Schema   string
	Port     int
}
