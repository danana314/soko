package store

import (
	"1008001/splitwiser/internal/models"
	"1008001/splitwiser/internal/utilities"
	_ "database/sql"
	"log/slog"
	"os"
	"time"

	"crawshaw.dev/jsonfile"
	_ "modernc.org/sqlite"
)

type Store struct {
	Trips []models.Trip
}

var db *jsonfile.JSONFile[Store]

func Init() {
	path := "./foo.db"
	var err error
	db, err = jsonfile.Load[Store](path)
	if os.IsNotExist(err) {
		db, err = jsonfile.New[Store](path)
		if err != nil {
			slog.Error(err.Error())
		}
		db.Write(func(s *Store) error {
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
			s.Trips = []models.Trip{
				{
					Id:        "test",
					Type:      models.TypeTrip,
					Users:     users,
					StartDate: startDate,
					EndDate:   endDate,
					Dates:     dates,
					Schedule:  make([]models.ScheduleEntry, len(users)*len(utilities.Range(startDate, endDate))),
				},
			}
			return nil
		})
	}

}

func GetTrip(tripId string) *models.Trip {
	var trip *models.Trip
	db.Read(func(s *Store) {
		for _, t := range s.Trips {
			if t.Id == tripId {
				trip = &t
			}
		}
	})
	return trip
}

func UpdateTrip(newTrip *models.Trip) *models.Trip {
	activeTrip := GetTrip(newTrip.Id)
	activeTrip.Name = newTrip.Name
	activeTrip.Type = newTrip.Type
	activeTrip.StartDate = newTrip.StartDate
	activeTrip.EndDate = newTrip.EndDate

	if activeTrip.StartDate != newTrip.StartDate || activeTrip.EndDate != newTrip.EndDate {
		slog.Info("dates updated. TODO - update list")
	}

	for ix, t := range inMemStore.Trips {
		if t.Id == activeTrip.Id {
			inMemStore.Trips[ix] = *activeTrip
		}
	}
	return activeTrip
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
