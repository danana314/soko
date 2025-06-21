package store

import (
	"1008001/soko/internal/safari/models"
	"time"
)

func SeedDB(store *DB) {
	// trips
	tripId := "test"
	trip := &models.Trip{
		Id:        tripId,
		Name:      "test",
		StartDate: time.Date(2024, time.January, 14, 0, 0, 0, 0, time.Local),
		EndDate:   time.Date(2024, time.February, 10, 0, 0, 0, 0, time.Local),
	}
	SaveTrip(trip)

	// users
	user1 := models.User{
		Id:   "testuser1",
		Name: "John Smith",
	}
	user2 := models.User{
		Id:   "testuser2",
		Name: "Jane Smith",
	}
	user3 := models.User{
		Id:   "testuser3",
		Name: "Will I Am",
	}
	AddUser(tripId, &user1)
	AddUser(tripId, &user2)
	AddUser(tripId, &user3)

	// expenses
	SaveExpense(tripId, &models.Expense{
		Id:           "testexpense1",
		Date:         time.Date(2024, time.January, 15, 0, 0, 0, 0, time.Local),
		Description:  "food at restaurant",
		Amount:       "20.34",
		PaidBy:       user1,
		Participants: []models.User{user1, user2},
	})

	SaveExpense(tripId, &models.Expense{
		Id:           "testexpense2",
		Date:         time.Date(2024, time.January, 21, 0, 0, 0, 0, time.Local),
		Description:  "food at another restaurant",
		Amount:       "20.34",
		PaidBy:       user2,
		Participants: []models.User{user1, user3},
	})
}
