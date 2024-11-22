package models

import (
	"time"
)

type Trip struct {
	Id        string
	Name      string `schema:"TripName"`
	Users     []User
	StartDate time.Time
	EndDate   time.Time
	Schedule  []ScheduleEntry
	Expenses  []Expense
}

func ValidateTrip(t *Trip) bool {
	// TODO
	return true
}
