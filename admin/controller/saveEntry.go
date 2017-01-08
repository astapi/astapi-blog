package controller

import (
	"fmt"
	"github.com/astapi/astapi-blog/model"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
	"strings"
	"time"
)

type NotFound struct {
}

func (e *NotFound) Error() string {
	return "Not Found"
}

func tags2Array(tags string) []string {
	if tags == "" {
		return make([]string, 0)
	}
	ret := strings.Split(tags, ",")
	return ret
}

func SaveEntry(w http.ResponseWriter, r *http.Request) {
	// bodyの値をパースする
	r.ParseForm()
	// パースされた値を取得する
	params := r.PostForm

	title := params.Get("title")
	body := params.Get("body")
	tags := tags2Array(params.Get("tags"))

	if title == "" || body == "" {
		http.Error(w, title, 500)
		return
	}

	// DataStoreに保存する
	ctx := appengine.NewContext(r)
	key := datastore.NewKey(ctx, "Entry", "", 0, nil)
	entry := &model.Entry{
		Title:     title,
		Body:      body,
		Tags:      tags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if _, err := datastore.Put(ctx, key, entry); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "ok")
}
