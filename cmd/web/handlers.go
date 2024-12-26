package main

import (
	"fmt"
	"html/template"
	// "log"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Check if the path is exactly '/'
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// Initailize a slice containing the path to the two files
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}

	// Parse the html template
	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)

		return

	}

	// Apply the parsed html template to w(the response)
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}

	// fmt.Fprint(w, "Hello from Snippetbox")
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"

	fmt.Fprint(w, "Create a new snippet")
}
