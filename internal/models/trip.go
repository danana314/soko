package models

import (
	"time"
)

type User struct {
	Id   int
	Name string
}

type Trip struct {
	Users   []User
	StartDt time.Time
	EndDt   time.Time
}
