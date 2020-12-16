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
	if err := homeView.Template.ExecuteTemplate(w,homeView.Layout, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/html")
	if err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil); err != nil {
		panic(err)
	}
}

func main() {

	homeView= view.NewView("bootstrap","view/home.gohtml")
	contactView = view.NewView("bootstrap","view/contact.gohtml")


	template.New(" hahaha")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)

}
