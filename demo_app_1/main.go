package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello")

	listHandler := func(w http.ResponseWriter, r *http.Request) {
		// Template object, that wraps a call function and returns a template object.
		tmpl := template.Must(template.ParseFiles("index.html"))
		// Defining context that we want to pass to the template
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Form Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		// This method passes thbe context into the page
		tmpl.Execute(w, films)
	}

	addBookHandler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) // dummy delay
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		// Getting the template, because we need the block element
		tmpl := template.Must(template.ParseFiles("index.html"))
		// Sending the data into the block element, which will create a new element, using the block by the name.
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	http.HandleFunc("/", listHandler)
	http.HandleFunc("/add-film/", addBookHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
