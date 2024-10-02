package store

import (
	"1008001/splitwiser/internal/models"
	"1008001/splitwiser/internal/utilities"
	_ "database/sql"
	"errors"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"

	"crawshaw.dev/jsonfile"
	"github.com/lithammer/shortuuid/v4"
	_ "modernc.org/sqlite"
)

type Store struct {
	Trips []models.Trip
}

var db *jsonfile.JSONFile[Store]

func Init() {
	path := filepath.Join(os.TempDir(), "db.json")
	var err error
	db, err = jsonfile.Load[Store](path)
	if errors.Is(err, fs.ErrNotExist) {
		db, err = jsonfile.New[Store](path)
		if err != nil {
			slog.Error(err.Error())
		}
		err = db.Write(func(s *Store) error {
			s.Trips = models.SeedTrip()
			return nil
		})
		if err != nil {
			slog.Error(err.Error())
		}
	} else if err != nil {
		slog.Error(err.Error())
	}
}

func CreateNewTrip() string {
	trip := new(models.Trip)
	trip.Id = shortuuid.New()
	db.Write(func(s *Store) error {
		s.Trips = append(s.Trips, *trip)
		return nil
	})
	return trip.Id
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

func UpdateTrip(updatedTrip *models.Trip) *models.Trip {
	activeTrip := GetTrip(updatedTrip.Id)
	activeTrip.Name = updatedTrip.Name
	//users
	activeTrip.StartDate = updatedTrip.StartDate
	activeTrip.EndDate = updatedTrip.EndDate
	//dates
	//schedule

	if activeTrip.StartDate != updatedTrip.StartDate || activeTrip.EndDate != updatedTrip.EndDate {
		slog.Info("dates updated. TODO - update list")
	}

	db.Write(func(s *Store) error {
		for ix, t := range s.Trips {
			if t.Id == activeTrip.Id {
				s.Trips[ix] = *activeTrip
			}
		}
		return nil
	})

	return activeTrip
}

func GetScheduleEntryList(entries []models.ScheduleEntry, date utilities.Date, user string) []models.ScheduleEntry {
	var resultList []models.ScheduleEntry
	for _, entry := range entries {
		if entry.Date == date && entry.User.Name == user {
			resultList = append(resultList, entry)
		}
	}
	return resultList
}
