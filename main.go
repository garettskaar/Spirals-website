package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type page struct {
	Title       string
	Description string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/menus", menus)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":80", nil)
}
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

func index(w http.ResponseWriter, _ *http.Request) {
	index := page{
		Title:       "Spirals - Pizza . Sushi",
		Description: "Spirals pizza happily serves authentic pizza and sushi to the Hiddens Springs community",
	}
	err := tpl.ExecuteTemplate(w, "index.gohtml", index)
	handleError(w, err)
}

func about(w http.ResponseWriter, _ *http.Request) {
	about := page{
		Title:       "Spirals About",
		Description: "Spirals is family owned and operated by team Quarles",
	}
	err := tpl.ExecuteTemplate(w, "about.gohtml", about)
	handleError(w, err)
}

func contact(w http.ResponseWriter, _ *http.Request) {
	contact := page{
		Title:       "Spirals Contact",
		Description: "Contact Spirals today for pick up, delivery or any questions about the services we offer.",
	}
	err := tpl.ExecuteTemplate(w, "contact.gohtml", contact)
	handleError(w, err)
}

func menus(w http.ResponseWriter, _ *http.Request) {
	menus := page{
		Title:       "Spirals Menu",
		Description: "Take a look at what delicous food is on Spirals Menu",
	}
	err := tpl.ExecuteTemplate(w, "menus.gohtml", menus)
	handleError(w, err)
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
