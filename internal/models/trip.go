package models

import (
	"time"
)

type User struct {
	Id   int
	Name string
}

type Trip struct {
	Id      string
	Users   []User
	StartDt time.Time
	EndDt   time.Time
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
