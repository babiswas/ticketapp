package model

type Tenant struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Domain  string `json:"domain"`
	Contact string `json:"contact"`
}
