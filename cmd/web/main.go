package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// command line option
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// To serve files a directory
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Handle requests to file path
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	infoLog.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
