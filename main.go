package main

import (
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

type Contact struct {
	Name  string
	Email string
}

func newContact(name, email string) Contact {
	return Contact{
		Name:  name,
		Email: email,
	}
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func (d *Data) hasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

func newData() Data {
	return Data{
		Contacts: []Contact{
			newContact("John", "jd@gmail.com"),
			newContact("Clara", "cd@gmail.com"),
		},
	}
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func newFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Data Data
	Form FormData
}

func newPage() Page {
	return Page{
		Data: newData(),
		Form: newFormData(),
	}
}

func main() {
	router := http.NewServeMux()

	page := newPage()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index", page)
	})

	router.HandleFunc("POST /contacts", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		email := r.FormValue("email")
		if page.Data.hasEmail(email) {
			formData := newFormData()
			formData.Values["name"] = name
			formData.Values["email"] = email
			formData.Errors["email"] = "Email already exists"

			// renderTemplate(w, http.StatusUnprocessableEntity, "form", formData)
			http.Error(w, "email already exists", http.StatusUnprocessableEntity)
			return
		}
		page.Data.Contacts = append(page.Data.Contacts, newContact(name, email))
		renderTemplate(w, "display", page.Data)
	})

	n := negroni.Classic() // default middleware: panic recovery, logger, static serving
	n.UseHandler(router)
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: n,
	}

	server.ListenAndServe()
}
