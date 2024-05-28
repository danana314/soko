package main

import (
	"1008001/splitwiser/internal/store"
	"1008001/splitwiser/internal/utilities"
	"fmt"
	"html/template"
	"net/http"

	"github.com/urfave/negroni"
)

func renderTemplate(t *template.Template, w http.ResponseWriter, tmpl string, data any) {
	err := t.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	router := http.NewServeMux()
	var templates = template.Must(template.ParseGlob("web/templates/*.tmpl"))
	store.Init()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(templates, w, "index", nil)
	})

	router.HandleFunc("POST /t/new", func(w http.ResponseWriter, r *http.Request) {
		id := utilities.NewId()
		http.Redirect(w, r, fmt.Sprintf("/t/%s", id), http.StatusSeeOther)
	})

	router.HandleFunc("GET /t/{tripId}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("tripId")
		trip := store.GetTrip(id)
		renderTemplate(templates, w, "trip", trip)
	})

	router.HandleFunc("POST /t/{tripId}", func(w http.ResponseWriter, r *http.Request) {
		startDate := r.FormValue("StartDate")
		endDate := r.FormValue("EndDate")
		fmt.Println(startDate)
		fmt.Println(endDate)
	})

	n := negroni.Classic() // default middleware: panic recovery, logger, static serving
	n.UseHandler(router)
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: n,
	}

	server.ListenAndServe()
}
