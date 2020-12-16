package main

import (

	"galleryApp_golang/view"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var (
	homeView    *view.View
	contactView *view.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/html")
	if err := homeView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/html")
	if err := contactView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {

	homeView= view.NewView("view/home.gohtml")
	contactView = view.NewView("view/contact.gohtml")


	template.New(" hahaha")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)

}
