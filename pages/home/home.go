package home

import (
	"net/http"
	"html/template"
)

type Page struct {
	Title string
}

func Home(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "Home"}
	t, _ := template.ParseFiles("templates/base.html", "templates/home.html")
	t.Execute(w, p)
}