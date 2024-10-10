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

func SaveTrip(trip *models.Trip) {
	upsertTripDetails := `INSERT INTO trips(ref, name, start_date, end_date)
		VALUES (?, ?, ?, ?)
		ON CONFLICT(ref) DO UPDATE SET
			name=excluded.name,
			start_date=excluded.start_date,
			end_date=excluded.end_date;`
	res, err := db_instance.Exec(upsertTripDetails, trip.Ref, trip.Name, trip.StartDate, trip.EndDate)
	if err != nil {
		slog.Error(err.Error())
	}
	if num_rows, _ := res.RowsAffected(); num_rows == 0 {
		slog.Error(fmt.Sprintf("saving trip did not result in updates: %#v", trip))
	}
}

func GetTrip(tripRef string) *models.Trip {
	queryStatement := `
		SELECT id, ref, name, start_date, end_date
		from trips
		where ref=?
	`
	var trip *models.Trip = new(models.Trip)
	err := db_instance.QueryRow(queryStatement, tripRef).Scan(&trip.Id, &trip.Ref, &trip.Name, &trip.StartDate, &trip.EndDate)
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info(fmt.Sprintf("%#v", trip))
	return trip
}
