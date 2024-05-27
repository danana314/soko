package models

import (
	"1008001/splitwiser/internal/utilities"
)

type User struct {
	Id   int
	Name string
}

const (
	TypeTrip   string = "Trip"
	TypeHome   string = "Home"
	TypeCouple string = "Couple"
	TypeOther  string = "Other"
)

type Trip struct {
	Id      string
	Type    string
	Users   []User
	StartDt utilities.Date
	EndDt   utilities.Date
}

// type Data struct {
// }

// func newData() Data {
// 	return Data{}
// }

// type FormData struct {
// 	Values map[string]string
// 	Errors map[string]string
// }

// func newFormData() FormData {
// 	return FormData{
// 		Values: make(map[string]string),
// 		Errors: make(map[string]string),
// 	}
// }

// type Page struct {
// 	Data Data
// 	Form FormData
// }

// func newPage() Page {
// 	return Page{
// 		Data: newData(),
// 		Form: newFormData(),
// 	}
// }
