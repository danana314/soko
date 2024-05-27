package utilities

import "time"

type Date time.Time

func NewDate(year int, month time.Month, day int) Date {
	return Date(time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
}

func (d Date) String() string {
	return time.Time(d).Format("2006/01/02")
}
