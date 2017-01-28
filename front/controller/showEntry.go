package controller

import (
	"github.com/astapi/astapi-blog/model"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"html/template"
	"net/http"
	"strconv"
)

type showEntry struct {
	Title string
	Entry *model.Entry
	DsKey *datastore.Key
}

func ShowEntry(w http.ResponseWriter, r *http.Request) {
	entry_id, err := strconv.Atoi(r.URL.Path[len("/entry/"):])
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	// DataStoreからEntryを取得する
	ctx := appengine.NewContext(r)
	k := datastore.NewKey(ctx, "Entry", "", int64(entry_id), nil)
	e := new(model.Entry)
	if err := datastore.Get(ctx, k, e); err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	funcMap := template.FuncMap{
		"safe": func(text string) template.HTML { return template.HTML(text) },
	}
	se := showEntry{
		Title: "One By One",
		Entry: e,
		DsKey: k,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t := template.Must(template.New("show_entry.html").Funcs(funcMap).ParseFiles("templates/show_entry.html", "templates/entry.html", "templates/entry_footer.html", "templates/header.html"))
	if err := t.ExecuteTemplate(w, "all", se); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
