package models

import (
	"1008001/splitwiser/internal/utilities"
)

type User struct {
	Id   int
	Name string
}

const (
	TypeTrip   string = "Trip"
	TypeHome   string = "Home"
	TypeCouple string = "Couple"
	TypeOther  string = "Other"
)

type Trip struct {
	Id        string
	Type      string
	Users     []User
	StartDate utilities.Date
	EndDate   utilities.Date
	Dates     []utilities.Date
	Schedule  []ScheduleEntry
}

type ScheduleEntry struct {
	Date   utilities.Date
	User   string
	Booked bool
}
