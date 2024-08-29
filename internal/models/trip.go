package models

import (
	"1008001/splitwiser/internal/utilities"
)

type User struct {
	Id   int
	Name string
}

type Trip struct {
	Id        string
	Name      string `schema:"TripName"`
	Users     []User
	StartDate utilities.Date
	EndDate   utilities.Date
	Schedule  []ScheduleEntry
}

type ScheduleEntry struct {
	Date   utilities.Date
	User   User
	Booked bool
}

func (t Trip) DateRange() []utilities.Date {
	return utilities.Range(t.StartDate, t.EndDate)
}

func (t Trip) UpdateSchedule() []ScheduleEntry {
	//todo: fill out
	return nil
}
