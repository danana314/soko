package store

import (
	"context"
	"database/sql"
	"log/slog"
	"time"
)

var seed_db = []string{insertExpenses}

func SeedDB(queries *Queries) {
	ctx := context.Background()

	// trips
	if _, err := queries.SaveTripDetails(ctx, SaveTripDetailsParams{
		TripID:    "test",
		Name:      sql.NullString{String: "test", Valid: true},
		StartDate: sql.NullTime{Time: time.Date(2024, time.January, 14, 0, 0, 0, 0, time.Local), Valid: true},
		EndDate:   sql.NullTime{Time: time.Date(2024, time.February, 10, 0, 0, 0, 0, time.Local), Valid: true},
	}); err != nil {
		slog.Error(err.Error())
	}

	// users
	if _, err := queries.AddUser(ctx, AddUserParams{
		UserID: "testuser1",
		TripID: "test",
		Name:   sql.NullString{String: "John Smith", Valid: true},
	}); err != nil {
		slog.Error(err.Error())
	}
	if _, err := queries.AddUser(ctx, AddUserParams{
		UserID: "testuser2",
		TripID: "test",
		Name:   sql.NullString{String: "Jane Smith", Valid: true},
	}); err != nil {
		slog.Error(err.Error())
	}
	if _, err := queries.AddUser(ctx, AddUserParams{
		UserID: "testuser3",
		TripID: "test",
		Name:   sql.NullString{String: "Will I Am", Valid: true},
	}); err != nil {
		slog.Error(err.Error())
	}

	// expenses
}

const insertExpenses string = `
	INSERT INTO expenses (tripId, date, description, amount, paidByUserId, participants)
	VALUES ('test', '2024-01-15', 'food at restaurant', 143.50, 'testuser3', '');
`
