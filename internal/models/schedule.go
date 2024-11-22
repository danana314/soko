package models

import (
	"fmt"
	"strings"
	"time"
)

type ScheduleEntry struct {
	Date time.Time
	User User
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

func (t *Trip) NewScheduleEntry(userDateString string) (*ScheduleEntry, error) {
	userId, date, err := splitUserDate(userDateString)
	if err != nil {
		return nil, err
	}
	scheduleEntry := new(ScheduleEntry)
	scheduleEntry.Date = date

	scheduleEntry.User = *new(User)
	scheduleEntry.User.Id = userId
	for _, u := range t.Users {
		if u.Id == userId {
			scheduleEntry.User.Name = u.Name
			break
		}
	}
	if scheduleEntry.User.Name == "" {
		return nil, fmt.Errorf("User ID %s not found in trip %#v", userId, t)
	}
	return scheduleEntry, nil
}

func (t *Trip) IsBooked(u User, d time.Time) bool {
	for _, se := range t.Schedule {
		if se.User.Id == u.Id && se.Date == d {
			return true
		}
	}
	return false
}
