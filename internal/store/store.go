package store

import (
	"1008001/splitwiser/internal/models"
	"1008001/splitwiser/internal/util"
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

// const defaultTimeout = 3 * time.Second

type Store interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

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

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(2 * time.Hour)

	db_instance = &DB{db}

	// create tables
	if _, err := db_instance.Exec(ddl); err != nil {
		slog.Error(err.Error())
	}
	slog.Info("db initialised")

	// seed db
	if seedDb {
		SeedDB(db_instance)
		slog.Info("db seeded")
	}

	return db_instance, nil
}

func AddOrUpdateTripDetails(trip *models.Trip) {
	upsertTripDetailsStatement := `INSERT INTO trips(trip_id, name, start_date, end_date)
	VALUES (?, ?, ?, ?)
	ON CONFLICT(trip_id) DO UPDATE SET
		name=excluded.name,
		start_date=excluded.start_date,
		end_date=excluded.end_date;`
	res, err := db_instance.Exec(upsertTripDetailsStatement, trip.Id, trip.Name, trip.StartDate, trip.EndDate)
	if err != nil {
		slog.Error(err.Error())
	}
	if num_rows, _ := res.RowsAffected(); num_rows == 0 {
		slog.Error(fmt.Sprintf("saving trip did not result in updates: %#v", trip))
	}
}

func AddUser(tripId string, user *models.User) {
	insertUserStatement := `INSERT INTO users(user_id, trip_id, name)
		VALUES (?, ?, ?);`
	res, err := db_instance.Exec(insertUserStatement, user.Id, tripId, user.Name)
	if err != nil {
		slog.Error(err.Error())
	}
	if num_rows, _ := res.RowsAffected(); num_rows == 0 {
		slog.Error(fmt.Sprintf("saving user did not result in insert: %#v", user))
	}
}

func DeleteUser(tripId string, userId string) {
	return
	// deleteScheduleEntriesStatement := `
	// 	DELETE FROM schedule
	// 	WHERE trip_id=? and user_id=?;
	// `

	deleteUserStatement := `DELETE FROM users(user_id, trip_id, name)
		VALUES (?, ?, ?);`
	res, err := db_instance.Exec(deleteUserStatement, user.Id, tripId, user.Name)
	if err != nil {
		slog.Error(err.Error())
	}
	if num_rows, _ := res.RowsAffected(); num_rows == 0 {
		slog.Error(fmt.Sprintf("saving user did not result in insert: %#v", user))
	}
}

func SaveSchedule(trip *models.Trip) {
	deleteStatement := `DELETE FROM schedule
		WHERE trip_id=?;`
	insertStatement := `INSERT INTO schedule(trip_id, user_id, date)
		VALUES (?, ?, ?);`

	tx, err := db_instance.Begin()
	if err != nil {
		slog.Error(err.Error())
	}

	_, err = db_instance.Exec(deleteStatement, trip.Id)
	if err != nil {
		tx.Rollback()
		slog.Error(err.Error())
		return
	}
	slog.Info(util.PrintStruct(trip.Schedule))
	for _, se := range trip.Schedule {
		res, err := db_instance.Exec(insertStatement, trip.Id, se.User.Id, se.Date)
		if err != nil {
			tx.Rollback()
			slog.Info(util.PrintStruct(se))
			slog.Error(err.Error())
			return
		}
		if num_rows, _ := res.RowsAffected(); num_rows == 0 {
			tx.Rollback()
			slog.Error(fmt.Sprintf("saving schedule did not result in updates: %#v", trip))
			return
		}
	}

	if err := tx.Commit(); err != nil {
		slog.Error(err.Error())
	}
}

func GetTrip(tripId string) *models.Trip {
	var trip *models.Trip = new(models.Trip)

	// get trip
	queryStatement := `
		SELECT trip_id, name, start_date, end_date
		FROM trips
		WHERE trip_id=?;`
	err := db_instance.QueryRow(queryStatement, tripId).Scan(&trip.Id, &trip.Name, &trip.StartDate, &trip.EndDate)
	if err != nil {
		slog.Error(err.Error())
	}

	// get users
	queryStatement = `
		SELECT user_id, name
		FROM users
		WHERE trip_id=?;`
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

	// get schedule
	queryStatement = `
		SELECT s.date, s.user_id, u.name
		FROM schedule s
			INNER JOIN users u on s.user_id = u.user_id
		WHERE s.trip_id=?;`
	rows, err = db_instance.Query(queryStatement, tripId)
	if err != nil {
		slog.Error(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var scheduleEntry models.ScheduleEntry
		if err := rows.Scan(&scheduleEntry.Date, &scheduleEntry.User.Id, &scheduleEntry.User.Name); err != nil {
			slog.Error(err.Error())
		}
		trip.Schedule = append(trip.Schedule, scheduleEntry)
	}
	if err := rows.Err(); err != nil {
		slog.Error(err.Error())
	}

	// slog.Info(fmt.Sprintf("%#v", trip))
	return trip
}
