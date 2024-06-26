package store

import (
	"1008001/splitwiser/internal/models"
	"1008001/splitwiser/internal/utilities"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"crawshaw.dev/jsonfile"
)

func TestGetTrip(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "testgettrip.json")
	var err error
	db, err = jsonfile.New[Store](path)
	if err != nil {
		t.Fatal(err)
	}

	startDate := utilities.NewDate(2024, time.January, 14)
	endDate := utilities.NewDate(2024, time.January, 18)
	dates := utilities.Range(startDate, endDate)
	want := models.Trip{
		Id: "test",
		Users: []models.User{
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
		},
		StartDate: startDate,
		EndDate:   endDate,
		Dates:     dates,
		Schedule:  make([]models.ScheduleEntry, 3*len(utilities.Range(startDate, endDate))),
	}
	sampleData := []models.Trip{want}

	db.Write(func(s *Store) error {
		s.Trips = sampleData
		return nil
	})

	got := GetTrip("test")
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("got %+v, want %+v", *got, want)
	}

	got = GetTrip("new trip")
	sampleData = append(sampleData, *got)
	if !reflect.DeepEqual(got, new(models.Trip)) {
		t.Errorf("got %+v, wanted new, empty trip", *got)
	}
	db.Read(func(data *Store) {
		if !reflect.DeepEqual(data.Trips, sampleData) {
			t.Errorf("new trip not added to db")
		}
	})

}
