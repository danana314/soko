package main

import (
	"1008001/splitwiser/internal/store"
	"1008001/splitwiser/internal/utilities"
	"fmt"
	"html/template"
	"net/http"

	"github.com/urfave/negroni"
)

var templates = template.Must(template.ParseGlob("web/templates/*.tmpl"))

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	router := http.NewServeMux()
	store.Init()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index", nil)
	})

	router.HandleFunc("POST /t/new", func(w http.ResponseWriter, r *http.Request) {
		id := utilities.NewId()
		http.Redirect(w, r, fmt.Sprintf("/t/%s", id), http.StatusSeeOther)
	})

	router.HandleFunc("GET /t/{tripId}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("tripId")
		trip := store.GetTrip(id)
		renderTemplate(w, "trip", trip)
	})

	n := negroni.Classic() // default middleware: panic recovery, logger, static serving
	n.UseHandler(router)
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: n,
	}

	server.ListenAndServe()
}
