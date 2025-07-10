package model

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Date    string `json:"date"`
}
