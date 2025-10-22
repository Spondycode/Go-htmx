package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

var films = []Film{
	{Title: "Inception", Director: "Christopher Nolan"},
	{Title: "The Matrix", Director: "The Wachowskis"},
	{Title: "Interstellar", Director: "Christopher Nolan"},
	{Title: "Parasite", Director: "Bong Joon-ho"},
}

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		data := map[string][]Film{
			"Films": films,
		}
		tmpl.Execute(w, data)
	}

	addFilm := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		title := r.FormValue("title")
		director := r.FormValue("director")

		if title != "" && director != "" {
			newFilm := Film{
				Title:    title,
				Director: director,
			}
			films = append(films, newFilm)

			tmpl := template.Must(template.ParseFiles("film-card.html"))
			tmpl.Execute(w, newFilm)
		}
	}

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film", addFilm)
	fmt.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
