package domain

type Passenger struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Document  string `json:"document"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
