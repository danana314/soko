package store

import (
	"1008001/splitwiser/internal/models"
	"1008001/splitwiser/internal/utilities"
	"time"
)

type Store struct {
	Trips []models.Trip
}

var inMemStore *Store

func Init() {
	users := []models.User{
		{
			Id:   1,
			Name: "John Smith",
		},
		{
			Id:   2,
			Name: "Jane Smith",
		},
		{
			Id:   3,
			Name: "Will I Am",
		},
	}
	startDate := utilities.NewDate(2024, time.January, 14)
	endDate := utilities.NewDate(2024, time.February, 10)
	dates := utilities.Range(startDate, endDate)
	inMemStore = &Store{
		Trips: []models.Trip{
			{
				Id:        "test",
				Type:      models.TypeTrip,
				Users:     users,
				StartDate: startDate,
				EndDate:   endDate,
				Dates:     dates,
				Schedule:  make([]models.ScheduleEntry, len(users)*len(utilities.Range(startDate, endDate))),
			},
		},
	}
}

func GetTrip(tripId string) *models.Trip {
	for _, t := range inMemStore.Trips {
		if t.Id == tripId {
			return &t
		}
	}
	return nil
}

func GetScheduleEntryList(entries []models.ScheduleEntry, date utilities.Date, user string) []models.ScheduleEntry {
	var resultList []models.ScheduleEntry
	for _, entry := range entries {
		if entry.Date == date && entry.User == user {
			resultList = append(resultList, entry)
		}
	}
	return resultList
}
