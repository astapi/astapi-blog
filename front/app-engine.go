package main

import (
	"controller"
	"net/http"
)

func init() {
	http.HandleFunc("/about/", controller.About)
	http.HandleFunc("/entry/", controller.ShowEntry)
	http.HandleFunc("/entry/tag/", controller.TagEntryList)
	http.HandleFunc("/", controller.Index)
}
