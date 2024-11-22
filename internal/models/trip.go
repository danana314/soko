package models

import (
	"1008001/splitwiser/internal/util"
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

func NewTrip() *Trip {
	return &Trip{
		Id: util.GenerateID(),
	}
}

func ValidateTrip(t *Trip) bool {
	// TODO
	return true
}
