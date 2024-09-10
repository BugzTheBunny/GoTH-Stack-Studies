package main

import (
	"html/template"
	"net/http"
	"strings"
	"time"
)

// Uppercase a string
func toUpper(str string) string {
	return strings.ToUpper(str)
}

// Format a date
func formatDate(t time.Time) string {
	return t.Format("January 2,2007")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(
			template.ParseGlob("templates/*.html"))

		data := struct {
			Name        string
			Title       string
			Description string
			Socials     map[string]string
			Features    []string
		}{
			Name:        "Tom",
			Title:       "Visitor",
			Description: "This paragraph is the new body of our page.",
			Socials: map[string]string{
				"Twitter/X": "@example",
				"LinkedIn":  "example-inc",
				"Instagram": "examplepics",
			},
			Features: []string{
				"Customizable Products",
				"24/7 Service",
				"Reliable and Secure",
			},
		}

		err := tmpl.ExecuteTemplate(w, "home.html", data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/functions", func(w http.ResponseWriter, r *http.Request) {

		funcMap := template.FuncMap{
			"toUpper":    toUpper,
			"formatDate": formatDate,
		}

		tmpl := template.Must(
			template.
				New("functions.html").
				Funcs(funcMap).
				ParseFiles("functions.html"))

		data := struct {
			Name        string
			CurrentDate time.Time
			Number      int
			Items       []string
		}{
			Name:        "John Doe",
			CurrentDate: time.Now(),
			Number:      7,
			Items: []string{
				"Apples", "Oranges", "Bananas",
			}}

		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":3000", nil)
}
