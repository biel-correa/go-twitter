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
	t, _ := template.ParseFiles("templates/home.html", "templates/header.html")
	t.Execute(w, p)
}