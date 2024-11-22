package models

import "1008001/splitwiser/internal/util"

type User struct {
	Id   string
	Name string
}

func NewUser() *User {
	user := &User{
		Id: util.GenerateID(),
	}
	return user
}

func ValidateUser(u *User) bool {
	// TODO
	return true
}
