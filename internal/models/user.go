package models

import "github.com/lithammer/shortuuid/v4"

type User struct {
	Id   string
	Name string
}

func ValidateUser(u *User) bool {
	// TODO
	return true
}

func NewUser(name string) *User {
	user := new(User)
	user.Id = shortuuid.New()
	user.Name = name
	return user
}
