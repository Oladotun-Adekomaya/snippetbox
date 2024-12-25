package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// To serve files a directory
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// Handle requests to file path
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux

}
