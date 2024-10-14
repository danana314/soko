package models

import (
	"1008001/splitwiser/internal/utilities"
	"time"

	"github.com/lithammer/shortuuid/v4"
)

type User struct {
	Id   string
	Name string
}

type Trip struct {
	Id        string
	Name      string `schema:"TripName"`
	Users     []User
	StartDate time.Time
	EndDate   time.Time
	Schedule  []ScheduleEntry
}

type ScheduleEntry struct {
	Date   utilities.Date
	User   User
	Booked bool
}

func NewTrip() *Trip {
	trip := new(Trip)
	trip.Id = shortuuid.New()
	return trip
}

func NewUser() *User {
	user := new(User)
	user.Id = shortuuid.New()
	return user
}

// func (t *Trip) DateRange() []utilities.Date {
// 	return utilities.Range(t.StartDate, t.EndDate)
// }

func (T *Trip) UserDateId(u User, d utilities.Date) string {
	return string(u.Id) + "_" + d.String()
}
