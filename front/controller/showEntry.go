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

var sEntryTmp = template.Must(template.New("show_entry.html").Funcs(funcMap).ParseFiles("templates/show_entry.html", "templates/entry.html", "templates/entry_footer.html", "templates/header.html"))

func ShowEntry(w http.ResponseWriter, r *http.Request) {
	eid, err := strconv.ParseInt(r.URL.Path[len("/entry/"):], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	// DataStoreからEntryを取得する
	ctx := appengine.NewContext(r)
	k := datastore.NewKey(ctx, "Entry", "", eid, nil)
	e := new(model.Entry)
	if err := datastore.Get(ctx, k, e); err != nil {
		http.Error(w, err.Error(), 500)
	}

	se := showEntry{
		Title: "One By One",
		Entry: e,
		DsKey: k,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := sEntryTmp.ExecuteTemplate(w, "all", se); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
