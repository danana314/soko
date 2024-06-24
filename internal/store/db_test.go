package store

import (
	"1008001/splitwiser/internal/models"
	"1008001/splitwiser/internal/utilities"
	"path/filepath"
	"testing"
	"time"

	"crawshaw.dev/jsonfile"
)

func TestGetTrip(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "testbasic.json")
	data, err := jsonfile.New[Store](path)
	if err != nil {
		t.Fatal(err)
	}

	data.Write(func(s *Store) error {
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
