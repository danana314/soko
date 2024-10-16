package db

import (
	"1008001/splitwiser/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const defaultTimeout = 3 * time.Second

type DB struct {
	*sql.DB
}

var db_instance *DB

func Init(dsn string) (*DB, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	// defer cancel()

	seedDb := false
	if _, err := os.Stat(dsn); errors.Is(err, os.ErrNotExist) {
		seedDb = true
	}

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(2 * time.Hour)

	for _, tbl := range init_db {
		if _, err := db.Exec(tbl); err != nil {
			slog.Error(err.Error())
		}
	}
	slog.Info("db initialised")

	if seedDb {
		for _, insStmt := range seed_db {
			if _, err := db.Exec(insStmt); err != nil {
				slog.Error(err.Error())
			}
		}
		slog.Info("db seeded")
	}

	db_instance = &DB{db}
	return db_instance, nil
}

func SaveTripDetails(trip *models.Trip) {
	upsertTripDetails := `INSERT INTO trips(tripId, name, startDate, endDate)
		VALUES (?, ?, ?, ?)
		ON CONFLICT(tripId) DO UPDATE SET
			name=excluded.name,
			startDate=excluded.startDate,
			endDate=excluded.endDate;`
	res, err := db_instance.Exec(upsertTripDetails, trip.Id, trip.Name, trip.StartDate, trip.EndDate)
	if err != nil {
		slog.Error(err.Error())
	}
	if num_rows, _ := res.RowsAffected(); num_rows == 0 {
		slog.Error(fmt.Sprintf("saving trip did not result in updates: %#v", trip))
	}
}

func AddUser(tripId string, user *models.User) {
	insertUser := `INSERT INTO users(userId, tripId, name)
		VALUES (?, ?, ?);`
	res, err := db_instance.Exec(insertUser, user.Id, tripId, user.Name)
	if err != nil {
		slog.Error(err.Error())
	}
	if num_rows, _ := res.RowsAffected(); num_rows == 0 {
		slog.Error(fmt.Sprintf("saving user did not result in insert: %#v", user))
	}
}

func GetTrip(tripId string) *models.Trip {
	var trip *models.Trip = new(models.Trip)

	// get trip
	queryStatement := `
		SELECT tripId, name, startDate, endDate
		FROM trips
		WHERE tripId=?;`
	err := db_instance.QueryRow(queryStatement, tripId).Scan(&trip.Id, &trip.Name, &trip.StartDate, &trip.EndDate)
	if err != nil {
		slog.Error(err.Error())
	}

	// get users
	queryStatement = `
		SELECT userId, name
		FROM users
		WHERE tripId=?;`
	rows, err := db_instance.Query(queryStatement, tripId)
	if err != nil {
		slog.Error(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			slog.Error(err.Error())
		}
		trip.Users = append(trip.Users, user)
	}
	if err := rows.Err(); err != nil {
		slog.Error(err.Error())
	}

	slog.Info(fmt.Sprintf("%#v", trip))
	return trip
}
