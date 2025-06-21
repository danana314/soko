package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

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
	if _, err := db_instance.Exec(schema); err != nil {
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

func exec(db *DB, query string, params ...any) error {
	res, err := db.Exec(query, params...)
	if err != nil {
		return err
	}
	if num_rows, _ := res.RowsAffected(); num_rows == 0 {
		return errors.New(fmt.Sprintln("Error executing query with params:", query, params))
	}
	return nil
}

// func execResult(db *DB, query string, params ...any) (sql.Result, error) {
// 	return nil, nil
// }
