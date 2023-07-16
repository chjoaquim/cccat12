package domain

type Driver struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Document  string `json:"document"`
	CarPlate  string `json:"car_plate"`
	CreatedAt string `json:"created_at"`
}
