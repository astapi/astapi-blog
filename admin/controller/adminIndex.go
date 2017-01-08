package controller

import (
	"github.com/astapi/astapi-blog/model"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"html/template"
	"net/http"
)

type Entry struct {
	Entry model.Entry
	DsKey *datastore.Key
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	query := datastore.NewQuery("Entry")

	var entries []model.Entry
	keys, err := query.Order("CreatedAt").GetAll(ctx, &entries)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	var es []Entry
	for i := 0; len(entries) > i; i++ {
		es = append(es, Entry{
			Entry: entries[i],
			DsKey: keys[i],
		})
	}

	funcMap := template.FuncMap{
		"safe": func(text string) template.HTML { return template.HTML(text) },
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t := template.Must(template.New("index.html").Funcs(funcMap).ParseFiles("templates/index.html"))
	if err := t.Execute(w, es); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
