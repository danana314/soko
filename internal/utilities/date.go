package utilities

import (
	"fmt"
	"time"
)

type Date string

func NewDate(dateStr string) Date {
	if _, err := time.Parse("2006-01-02", dateStr); err != nil {
		fmt.Println(err.Error())
	}
	return Date(dateStr)
}

func Range(d1 Date, d2 Date) []Date {
	var startDate, endDate time.Time
	d1dt, _ := time.Parse("2006-01-02", string(d1))
	d2dt, _ := time.Parse("2006-01-02", string(d2))
	if d1dt.Before(d2dt) {
		startDate = d1dt
		endDate = d2dt
	} else {
		startDate = d2dt
		endDate = d1dt
	}
	diff := int(endDate.Sub(startDate).Hours()/24) + 1
	dates := make([]Date, diff)
	for i := 0; i < diff; i++ {
		nextDate := startDate.AddDate(0, 0, i)
		dates[i] = Date(nextDate.String())
	}
	return dates
}
