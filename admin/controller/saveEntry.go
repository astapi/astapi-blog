package controller

import (
	"github.com/astapi/astapi-blog/model"
	"golang.org/x/net/context"
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
	createdAt, err := initCreatedAt(params.Get("createdAt"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if title == "" || body == "" {
		http.Error(w, title, 500)
		return
	}

	// DataStoreに保存する
	ctx := appengine.NewContext(r)
	if err := saveTag(ctx, tags); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	key := datastore.NewKey(ctx, "Entry", "", 0, nil)
	entry := &model.Entry{
		Title:     title,
		Body:      body,
		Tags:      tags,
		CreatedAt: createdAt,
		UpdatedAt: time.Now(),
	}
	if _, err := datastore.Put(ctx, key, entry); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	http.Redirect(w, r, "/admin/", 303)
}

func initCreatedAt(param string) (time.Time, error) {
	if param == "" {
		return time.Now(), nil
	}
	const format = "2006/01/02 15:04"
	ret, err := time.Parse(format, param)
	return ret, err
}

func saveTag(c context.Context, tags []string) (err error) {
	if len(tags) == 0 {
		return
	}

	for _, v := range tags {
		key := datastore.NewKey(c, "Tag", v, 0, nil)
		tag := new(model.Tag)
		err := datastore.Get(c, key, tag)
		if err == nil {
			tag.EntryCount += 1
			tag.UpdatedAt = time.Now()
			if _, err := datastore.Put(c, key, tag); err != nil {
				return err
			}
			continue
		}

		ntag := &model.Tag{
			TagName:    v,
			EntryCount: 1,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		if _, err := datastore.Put(c, key, ntag); err != nil {
			return err
		}
	}
	return
}
