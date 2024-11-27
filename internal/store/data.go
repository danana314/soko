package store

import (
	"1008001/splitwiser/internal/models"
	"1008001/splitwiser/internal/util"
	_ "embed"
	"fmt"
	"log/slog"
)

//go:embed schema.sql
var schema string

func GetTrip(tripId string) *models.Trip {
	queryStatement := `
		SELECT trip_id, name, start_date, end_date
		FROM trips
		WHERE trip_id=?;`
	var tripDetail *models.Trip = &models.Trip{}

	err := db_instance.QueryRow(queryStatement, tripId).Scan(&tripDetail.Id, &tripDetail.Name, &tripDetail.StartDate, &tripDetail.EndDate)
	if err != nil {
		slog.Error(err.Error())
	}

	return tripDetail
}

func SaveTrip(trip *models.Trip) {
	upsertTripDetailsStatement := `
		INSERT INTO trips(trip_id, name, start_date, end_date)
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

func GetUsers(tripId string) []*models.User {
	queryStatement := `
		SELECT user_id, name
		FROM users
		WHERE trip_id=?;`
	users := make([]*models.User, 0)

	rows, err := db_instance.Query(queryStatement, tripId)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			slog.Error(err.Error())
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		slog.Error(err.Error())
		return nil
	}

	return users
}

func AddUser(tripId string, user *models.User) {
	insertUserStatement := `
		INSERT INTO users(user_id, trip_id, name)
		VALUES (?, ?, ?);`
	res, err := db_instance.Exec(insertUserStatement, user.Id, tripId, user.Name)
	if err != nil {
		slog.Error(err.Error())
	}
	if num_rows, _ := res.RowsAffected(); num_rows == 0 {
		slog.Error(fmt.Sprintf("saving user did not result in insert: %#v", user))
	}
}

func GetSchedule(tripId string) []*models.ScheduleEntry {
	queryStatement := `
		SELECT s.date, s.user_id, u.name
		FROM schedule s
			INNER JOIN users u on s.user_id = u.user_id
		WHERE s.trip_id=?;`
	schedule := make([]*models.ScheduleEntry, 0)

	rows, err := db_instance.Query(queryStatement, tripId)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		var scheduleEntry models.ScheduleEntry
		if err := rows.Scan(&scheduleEntry.Date, &scheduleEntry.User.Id, &scheduleEntry.User.Name); err != nil {
			slog.Error(err.Error())
		}
		schedule = append(schedule, &scheduleEntry)
	}
	if err := rows.Err(); err != nil {
		slog.Error(err.Error())
		return nil
	}

	return schedule
}

func SaveSchedule(tripId string, schedule []*models.ScheduleEntry) {
	deleteStatement := `DELETE FROM schedule
		WHERE trip_id=?;`
	insertStatement := `INSERT INTO schedule(trip_id, user_id, date)
		VALUES (?, ?, ?);`

	tx, err := db_instance.Begin()
	if err != nil {
		slog.Error(err.Error())
	}

	_, err = db_instance.Exec(deleteStatement, tripId)
	if err != nil {
		tx.Rollback()
		slog.Error(err.Error())
		return
	}
	for _, se := range schedule {
		res, err := db_instance.Exec(insertStatement, tripId, se.User.Id, se.Date)
		if err != nil {
			tx.Rollback()
			slog.Info(util.PrintStruct(se))
			slog.Error(err.Error())
			return
		}
		if num_rows, _ := res.RowsAffected(); num_rows == 0 {
			tx.Rollback()
			slog.Error(fmt.Sprintf("saving schedule did not result in updates: %#v", schedule))
			return
		}
	}

	if err := tx.Commit(); err != nil {
		slog.Error(err.Error())
	}
}

func GetExpenses(tripId string) []*models.Expense {
	getExpensesStatement := `
		SELECT
			expense_id, date, description, amount, paid_by_user_id, participants
		FROM expenses
		WHERE trip_id=?;`
	expenses := make([]*models.Expense, 0)

	rows, err := db_instance.Query(getExpensesStatement, tripId)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		var expense models.Expense
		if err := rows.Scan(&expense.Id, &expense.Date, &expense.Description, &expense.Amount, &expense.PaidBy.Id, &expense.Participants); err != nil {
			slog.Error(err.Error())
		}
		expenses = append(expenses, &expense)
	}
	if err := rows.Err(); err != nil {
		slog.Error(err.Error())
		return nil
	}

	return expenses
}

func SaveExpense(tripId string, expense *models.Expense) {
	insertExpenseStatement := `
		INSERT INTO expenses(expense_id, trip_id, date, description, amount, paid_by_user_id, participants)
		VALUES (?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(expense_id) DO UPDATE SET
			date=excluded.date,
			description=excluded.description,
			amount=excluded.amount,
			paid_by_user=excluded.paid_by_user,
			participants=excluded.participants;
	`
	err := exec(db_instance, insertExpenseStatement, expense.Id, tripId, expense.Date, expense.Description, expense.Amount, expense.PaidBy.Id, expense.Participants)
	if err != nil {
		slog.Error(err.Error())
	}
}

func GetTripData(tripId string) *models.TripData {
	tripDetails := GetTrip(tripId)
	users := GetUsers(tripId)
	schedule := GetSchedule(tripId)
	expenses := GetExpenses(tripId)

	return &models.TripData{
		Trip:     tripDetails,
		Users:    users,
		Schedule: schedule,
		Expenses: expenses,
	}
}
