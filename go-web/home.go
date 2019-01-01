package main

import (
	"net/http"
	"path"
	"text/template"
)

type Message struct {
	TextMessage string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := Message{"Building Web Apps with Go"}

		fp := path.Join("templates", "home.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, message); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8888", nil)
}
