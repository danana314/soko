package main

import (
	"1008001/splitwiser/internal/models"
	"1008001/splitwiser/internal/store"
	"1008001/splitwiser/internal/utilities"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/urfave/negroni"
)

var decoder = schema.NewDecoder()

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

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(templates, w, "index", nil)
	})

	router.HandleFunc("POST /t/new", func(w http.ResponseWriter, r *http.Request) {
		id := utilities.NewId()
		//todo: create new trip here

		http.Redirect(w, r, fmt.Sprintf("/t/%s", id), http.StatusSeeOther)
	})

	router.HandleFunc("GET /t/{tripId}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("tripId")
		trip := store.GetTrip(id)
		//todo: return 'trip not found' on nil
		renderTemplate(templates, w, "trip", trip)
	})

	router.HandleFunc("POST /t/{tripId}", func(w http.ResponseWriter, r *http.Request) {
		tripId := r.PathValue("tripId")
		err := r.ParseForm()
		if err != nil {
			slog.Error(err.Error())
		}

		trip := &models.Trip{}
		trip.Id = tripId
		err = decoder.Decode(trip, r.PostForm)
		if err != nil {
			slog.Error(err.Error(), "postform", r.PostForm)
		}
		trip = store.UpdateTrip(trip)
		renderTemplate(templates, w, "trip_detail", trip)
	})

	n := negroni.Classic() // default middleware: panic recovery, logger, static serving
	n.UseHandler(router)
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: n,
	}

	server.ListenAndServe()
}
