package store

import (
	"1008001/splitwiser/internal/models"
	"time"
)

type Store struct {
	Users []models.User
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
	inMemStore = &Store{
		Users: users,
		Trips: []models.Trip{
			{
				Id:      "gRvBazHJZGSXeFqsjABtBy",
				Users:   users,
				StartDt: time.Date(2024, time.January, 14, 0, 0, 0, 0, time.UTC),
				EndDt:   time.Date(2024, time.February, 10, 0, 0, 0, 0, time.UTC),
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
