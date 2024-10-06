package main

import (
	"1008001/splitwiser/assets"
	"1008001/splitwiser/internal/models"
	"1008001/splitwiser/internal/store"
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

func routes() http.Handler {
	router := http.NewServeMux()
	var templates = template.Must(template.ParseGlob("assets/templates/*.tmpl"))

	fileServer := http.FileServer(http.FS(assets.EmbeddedFiles))
	router.Handle("GET /static/", fileServer)

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(templates, w, "index", nil)
	})

	router.HandleFunc("POST /t/new", func(w http.ResponseWriter, r *http.Request) {
		id := store.CreateNewTrip()
		http.Redirect(w, r, fmt.Sprintf("/t/%s", id), http.StatusSeeOther)
	})

	router.HandleFunc("GET /t/{tripId}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("tripId")
		trip := store.GetTrip(id)
		if trip == nil {
			http.Error(w, "trip not found", http.StatusNotFound)
		} else {
			renderTemplate(templates, w, "trip", trip)
		}
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
		trip.UpdateTripDetails(trip)
		store.UpdateTrip(trip)
		renderTemplate(templates, w, "trip_detail", trip)
	})

	router.HandleFunc("POST /t/{tripId}/schedule", func(w http.ResponseWriter, r *http.Request) {
		tripId := r.PathValue("tripId")
		err := r.ParseForm()
		if err != nil {
			slog.Error(err.Error())
		}
		slog.Info(tripId)
		slog.Info(fmt.Sprintf("%#v", r.PostForm))
	})

	n := negroni.Classic() // default middleware: panic recovery, logger, static serving
	n.UseHandler(router)

	return n
}
