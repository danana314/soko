package models

import (
	"1008001/splitwiser/internal/util"
	"fmt"
	"strings"
	"time"
)

type TripData struct {
	Trip     *Trip
	Users    []*User
	Schedule []*ScheduleEntry
	Expenses []*Expense
}

type Trip struct {
	Id        string
	Name      string `schema:"TripName"`
	StartDate time.Time
	EndDate   time.Time
}

type User struct {
	Id   string
	Name string
}

type ScheduleEntry struct {
	Date time.Time
	User User
}

type Expense struct {
	Id           string
	Date         time.Time
	Description  string
	Amount       int64
	PaidBy       User
	Participants []User
}

func NewTrip() *Trip {
	return &Trip{
		Id: util.GenerateID(),
	}
}

func NewUser() *User {
	user := &User{
		Id: util.GenerateID(),
	}
	return user
}

func NewExpense() *Expense {
	return &Expense{
		Id: util.GenerateID(),
	}
}

func splitUserDate(userDateString string) (string, time.Time, error) {
	segments := strings.Split(userDateString, "_")
	if len(segments) != 2 {
		return "", time.Time{}, fmt.Errorf("Unexpected user-date string: " + userDateString)
	}
	userId := segments[0]
	date, err := time.Parse("2006-01-02", segments[1])
	if err != nil {
		return "", time.Time{}, err
	}
	return userId, date, nil
}

func NewScheduleEntry(users []*User, userDateString string) (*ScheduleEntry, error) {
	userId, date, err := splitUserDate(userDateString)
	if err != nil {
		return nil, err
	}
	scheduleEntry := new(ScheduleEntry)
	scheduleEntry.Date = date

	scheduleEntry.User = *new(User)
	scheduleEntry.User.Id = userId
	for _, u := range users {
		if u.Id == userId {
			scheduleEntry.User.Name = u.Name
			break
		}
	}
	if scheduleEntry.User.Name == "" {
		return nil, fmt.Errorf("User ID %s not found in users %#v", userId, users)
	}
	return scheduleEntry, nil
}

func (t *TripData) IsBooked(u User, d time.Time) bool {
	for _, se := range t.Schedule {
		if se.User.Id == u.Id && se.Date == d {
			return true
		}
	}
	return false
}
