package main

import (
	"1008001/splitwiser/internal/utilities"
	"fmt"
	"html/template"
	"net/http"

	"github.com/urfave/negroni"
)

var templates = template.Must(template.ParseGlob("templates/*.tmpl"))

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index", nil)
	})

	router.HandleFunc("POST /t/new", func(w http.ResponseWriter, r *http.Request) {
		id := utilities.NewId()
		http.Redirect(w, r, fmt.Sprintf("/t/%s", id), http.StatusSeeOther)
	})

	router.HandleFunc("GET /t/{tripId}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("tripId")
		fmt.Println(id)
		renderTemplate(w, "trip", id)
	})

	n := negroni.Classic() // default middleware: panic recovery, logger, static serving
	n.UseHandler(router)
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: n,
	}

	server.ListenAndServe()
}
