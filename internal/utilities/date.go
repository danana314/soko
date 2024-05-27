package utilities

import (
	"time"
)

type Date struct {
	time.Time
}

func NewDate(year int, month time.Month, day int) Date {
	d := Date{}
	d.Time = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return d
}

func (d Date) String() string {
	return d.Format("2006-01-02")
}

func Range(d1 Date, d2 Date) []Date {
	var startDate, endDate Date
	if d1.Before(d2.Time) {
		startDate = d1
		endDate = d2
	} else {
		startDate = d2
		endDate = d1
	}
	diff := int(endDate.Sub(startDate.Time).Hours()/24) + 1
	dates := make([]Date, diff)
	for i := 0; i < diff; i++ {
		nextDate := Date{}
		nextDate.Time = startDate.AddDate(0, 0, i)
		dates[i] = nextDate
	}
	return dates
}
