package models

import "time"

type Expense struct {
	Date         time.Time
	Description  string
	Amount       int64
	PaidBy       User
	Participants []User
}
