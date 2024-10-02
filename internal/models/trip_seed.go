package models

import (
	"1008001/splitwiser/internal/utilities"
	"time"
)

func SeedTrip() []Trip {
	startDate := utilities.NewDate(2024, time.January, 14)
	endDate := utilities.NewDate(2024, time.February, 10)
	trip := NewTrip()
	trip.Id = "test"
	trip.UpdateTripDetails(&Trip{Name: "test", StartDate: startDate, EndDate: endDate})
	trip.AddUser("John Smith")
	trip.AddUser("Jane Smith")
	trip.AddUser("Will I Am")
	return []Trip{*trip}
}
