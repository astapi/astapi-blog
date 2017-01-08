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

func Index(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	query := ds.NewQuery("Entry")

	var entries []model.Entry
	keys, err := query.Order("CreatedAt").GetAll(ctx, &entries)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	//[/Entry,4889528208719872 /Entry,6015428115562496 /Entry,6296903092273152]

	var es []Entry
	for i := 0; len(entries) > i; i++ {
		//	entries[i].ToHtml()
		es = append(es, Entry{
			Entry: entries[i],
			DsKey: keys[i],
		})
	}

	index := Indexa{
		Title:   "あすたぴのブログ",
		Entries: es,
	}

	funcMap := template.FuncMap{
		"safe": func(text string) template.HTML { return template.HTML(text) },
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t := template.Must(template.New("index.html").Funcs(funcMap).ParseFiles("templates/index.html", "templates/entry.html", "templates/entry_footer.html", "templates/header.html"))
	if err := t.ExecuteTemplate(w, "all", index); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
