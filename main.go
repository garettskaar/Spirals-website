package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/menus", menus)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	handleError(w, err)
}

func about(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	handleError(w, err)
}

func contact(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	handleError(w, err)
}

func menus(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "menus.gohtml", nil)
	handleError(w, err)
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
