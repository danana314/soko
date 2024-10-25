package main

import (
	"1008001/splitwiser/internal/db"
	"1008001/splitwiser/internal/funcs"
	"1008001/splitwiser/internal/models"
	"1008001/splitwiser/web"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"reflect"
	"time"

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

var timeConverter = func(value string) reflect.Value {
	if v, err := time.Parse("2006-01-02", value); err == nil {
		return reflect.ValueOf(v)
	}
	return reflect.Value{} // this is the same as the private const invalidType
}

func routes() http.Handler {
	router := http.NewServeMux()
	decoder.RegisterConverter(time.Time{}, timeConverter)

	fileServer := http.FileServer(http.FS(web.EmbeddedFiles))
	router.Handle("GET /static/", fileServer)

	// var templates = template.Must(template.ParseGlob("web/templates/*.tmpl")).Funcs(funcs.TemplateFuncs)
	templates, err := template.New("").Funcs(funcs.TemplateFuncs).ParseFS(web.EmbeddedFiles, "templates/*.tmpl")
	if err != nil {
		slog.Error(err.Error())
	}
	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(templates, w, "index", nil)
	})

	router.HandleFunc("POST /t/new", func(w http.ResponseWriter, r *http.Request) {
		trip := models.NewTrip()
		db.SaveTripDetails(trip)
		http.Redirect(w, r, fmt.Sprintf("/t/%s", trip.Id), http.StatusSeeOther)
	})

	router.HandleFunc("GET /t/{tripId}", func(w http.ResponseWriter, r *http.Request) {
		tripId := r.PathValue("tripId")
		trip := db.GetTrip(tripId)
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

		trip := new(models.Trip)
		trip.Id = tripId
		err = decoder.Decode(trip, r.PostForm)
		if err != nil {
			slog.Error(err.Error(), "postform", r.PostForm)
		}
		db.SaveTripDetails(trip)
		trip = db.GetTrip(trip.Id)
		renderTemplate(templates, w, "trip_detail", trip)
	})

	router.HandleFunc("POST /t/{tripId}/u", func(w http.ResponseWriter, r *http.Request) {
		tripId := r.PathValue("tripId")
		trip := db.GetTrip(tripId)

		name := r.PostFormValue("name")
		user := models.NewUser(name)
		db.AddUser(tripId, user)
		trip.Users = append(trip.Users, *user)
		renderTemplate(templates, w, "trip_detail", trip)
	})

	router.HandleFunc("POST /t/{tripId}/s", func(w http.ResponseWriter, r *http.Request) {
		tripId := r.PathValue("tripId")
		err := r.ParseForm()
		if err != nil {
			slog.Error(err.Error())
		}
		trip := db.GetTrip(tripId)
		for k, _ := range r.PostForm {
			se, err := trip.NewScheduleEntry(k)
			if err != nil {
				slog.Error(err.Error())
			}
			trip.Schedule = append(trip.Schedule, *se)
		}
		db.SaveSchedule(trip)
		renderTemplate(templates, w, "trip_detail", trip)
	})

	router.HandleFunc("POST /t/{tripId}/e", func(w http.ResponseWriter, r *http.Request) {
		tripId := r.PathValue("tripId")
		err := r.ParseForm()
		if err != nil {
			slog.Error(err.Error())
		}
		expense := new(models.Expense)
		err = decoder.Decode(expense, r.PostForm)
		if err != nil {
			slog.Error(err.Error(), "postform", r.PostForm)
		}
		trip := db.GetTrip(tripId)
		trip.Expenses = append(trip.Expenses, *expense)
		slog.Info(fmt.Sprintf("%#v", trip))
		renderTemplate(templates, w, "trip_detail", trip)
	})

	n := negroni.Classic() // default middleware: panic recovery, logger, static serving
	n.UseHandler(router)

	return n
}
