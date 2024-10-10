package models

import (
	"1008001/splitwiser/internal/utilities"
	"time"

	"github.com/lithammer/shortuuid/v4"
)

type User struct {
	Id   int
	Name string
}

type Trip struct {
	Id        int
	Ref       string
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

// func (t *Trip) DateRange() []utilities.Date {
// 	return utilities.Range(t.StartDate, t.EndDate)
// }

func (T *Trip) UserDateId(u User, d utilities.Date) string {
	return string(u.Id) + "_" + d.String()
}

// func (t *Trip) initSchedule() {
// 	t.Schedule = make(map[ScheduleKey]bool)
// 	for _, d := range utilities.Range(t.StartDate, t.EndDate) {
// 		for _, u := range t.Users {
// 			t.Schedule[ScheduleKey{Date: d, User: u}] = false
// 		}
// 	}
// }

func NewTrip() *Trip {
	trip := new(Trip)
	trip.Ref = shortuuid.New()

	return trip
}

func (t *Trip) UpdateTripDetails(updatedTrip *Trip) {
	t.Name = updatedTrip.Name
	t.StartDate = updatedTrip.StartDate
	t.EndDate = updatedTrip.EndDate

	if t.StartDate != updatedTrip.StartDate || t.EndDate != updatedTrip.EndDate {
		// if t.Schedule == nil {
		// 	t.initSchedule()
		// } else {
		// 	if t.StartDate.Before(updatedTrip.StartDate.Time) {
		// 		// trim
		// 		for k := range t.Schedule {
		// 			if k.Date.Before(updatedTrip.StartDate.Time) {
		// 				delete(t.Schedule, k)
		// 			}
		// 		}
		// 	} else if t.StartDate.After(updatedTrip.StartDate.Time) {
		// 		// add
		// 		for _, d := range utilities.Range(updatedTrip.StartDate, t.StartDate) {
		// 			for _, u := range t.Users {
		// 				t.Schedule[ScheduleKey{Date: d, User: u}] = false
		// 			}
		// 		}
		// 	}

		// 	if t.EndDate.Before(updatedTrip.EndDate.Time) {
		// 		// add
		// 		for _, d := range utilities.Range(t.EndDate, updatedTrip.EndDate) {
		// 			for _, u := range t.Users {
		// 				t.Schedule[ScheduleKey{Date: d, User: u}] = false
		// 			}
		// 		}
		// 	} else if t.EndDate.After(updatedTrip.EndDate.Time) {
		// 		// trim
		// 		for k := range t.Schedule {
		// 			if k.Date.After(updatedTrip.EndDate.Time) {
		// 				delete(t.Schedule, k)
		// 			}
		// 		}
		// 	}
		// }

		// t.StartDate = updatedTrip.StartDate
		// t.EndDate = updatedTrip.EndDate
	}
}

// func (t *Trip) AddUser(newUserName string) {
// 	user := User{Id: shortuuid.New(), Name: newUserName}
// 	t.Users = append(t.Users, user)
// 	// for _, d := range utilities.Range(t.StartDate, t.EndDate) {
// 	// 	t.Schedule[ScheduleKey{Date: d, User: user}] = false
// 	// }
// }

// func (t *Trip) DeleteUser(userId string) {
// 	for _, u := range t.Users {
// 		if u.Id == userId {
// 			// TODO delete from array
// 			u = User{}
// 		}
// 	}
// 	// for k := range t.Schedule {
// 	// 	if k.User.Id == userId {
// 	// 		delete(t.Schedule, k)
// 	// 	}
// 	// }
// }
