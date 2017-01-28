package controller

import (
	"html/template"
	"net/http"
)

var about = template.Must(template.ParseFiles("templates/about.html"))

func About(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := about.Execute(w, ""); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
