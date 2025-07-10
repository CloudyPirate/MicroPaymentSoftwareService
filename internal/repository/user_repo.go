package repository

import "github.com/CloudyPirate/MPSS/internal/model"

var users []model.User

func GetUsers() []model.User {
	return users
}

func AddUser(user model.User) {
	users = append(users, user)
}
