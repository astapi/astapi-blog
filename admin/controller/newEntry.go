package controller

import (
	"html/template"
	"net/http"
)

func NewEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t := template.Must(template.New("new_entry.html").ParseFiles("templates/new_entry.html"))
	if err := t.Execute(w, ""); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
