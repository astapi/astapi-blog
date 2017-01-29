package controller

import (
	"github.com/astapi/astapi-blog/model"
	"google.golang.org/appengine"
	ds "google.golang.org/appengine/datastore"
	"html/template"
	"net/http"
)

type Entry struct {
	Entry model.Entry
	DsKey *ds.Key
}

type Indexa struct {
	Title   string
	Entries []Entry
}

var funcMap = template.FuncMap{
	"safe": func(text string) template.HTML { return template.HTML(text) },
}
var indexTemp = template.Must(template.New("index.html").Funcs(funcMap).ParseFiles("templates/index.html", "templates/entry.html", "templates/entry_footer.html", "templates/header.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	query := ds.NewQuery("Entry")

	var entries []model.Entry
	keys, err := query.Order("-CreatedAt").GetAll(ctx, &entries)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	es := make([]Entry, 0, len(entries))
	for i := 0; len(entries) > i; i++ {
		es = append(es, Entry{
			Entry: entries[i],
			DsKey: keys[i],
		})
	}

	index := Indexa{
		Title:   "One By One",
		Entries: es,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := indexTemp.ExecuteTemplate(w, "all", index); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
