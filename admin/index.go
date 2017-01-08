package main

import (
	"controller"
	"net/http"
)

func init() {
	http.HandleFunc("/admin/entry/new/", controller.NewEntry)
	http.HandleFunc("/admin/entry/save/", controller.SaveEntry)
	http.HandleFunc("/admin/", controller.AdminIndex)
}
