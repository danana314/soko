package store

import (
	"1008001/splitwiser/internal/models"
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
	AddOrUpdateTripDetails(trip)

	// users
	AddUser(tripId, &models.User{
		Id:   "testuser1",
		Name: "John Smith",
	})
	AddUser(tripId, &models.User{
		Id:   "testuser2",
		Name: "Jane Smith",
	})
	AddUser(tripId, &models.User{
		Id:   "testuser3",
		Name: "Will I Am",
	})

	// expenses

	// if _, err := queries.AddExpense(ctx, AddExpenseParams{
	// 	TripID:       "test",
	// 	Date:         sql.NullTime{Time: time.Date(2024, time.January, 15, 0, 0, 0, 0, time.Local), Valid: true},
	// 	Description:  sql.NullString{String: "food at restaurant", Valid: true},
	// 	Amount:       sql.NullFloat64{Float64: 143.50, Valid: true},
	// 	PaidByUserID: sql.NullString{String: "testuser3", Valid: true},
	// 	Participants: []byte("[testuser1]"),
	// }); err != nil {
	// 	slog.Error(err.Error())
	// }
}
